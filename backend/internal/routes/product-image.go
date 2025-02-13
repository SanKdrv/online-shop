package routes

import (
	"backend/internal/lib/api/response"
	"backend/internal/types"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// @Summary getImageByHash
// @Tags Product Image
// @Description get image file by image hashName
// @ID get-image-by-hash
// @Accept json
// @Produce image/jpeg
// @Param hash_name query string true "Название изображения (Хеш)"
// @Success 200 {file} binary "Image file"
// @Failure 400 {object} string "Ошибка валидации (например, отсутствует hash_name)"
// @Failure 404 {object} string "Изображение не найдено"
// @Failure 500 {object} string "Внутренняя ошибка сервера"
// @Router /api/product-image/get-image-by-hash [get]
func (h *Handler) getImageByHash(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product-image.getImageByHash"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		hashName := r.URL.Query().Get("hash_name")
		if hashName == "" {
			log.Error("missing hash_name", slog.String("error", "missing hash_name"))
			return
		}

		imagePath := "uploads/" + hashName
		if _, err := os.Stat(imagePath); os.IsNotExist(err) {
			log.Error("file not found", slog.String("hash_name", hashName))
			http.Error(w, "File not found", http.StatusNotFound)
			//w.Write()
			return
		}

		// Определяем Content-Type
		w.Header().Set("Content-Type", "image/jpeg") // Можно улучшить, используя `http.DetectContentType`

		// Отправляем файл
		http.ServeFile(w, r, imagePath)
	}
}

// @Summary getImagesByProductId
// @Tags Product Image
// @Description getting images by product id
// @ID get-images-by-product-id
// @Accept  json
// @Produce  json
// @Param product_id query int64 true "ID продукта"
// @Success 200 {object} types.GetImageHashByProductIdResponse
// @Failure 400,404 {object} types.GetImageHashByProductIdResponse
// @Failure 500 {object} types.GetImageHashByProductIdResponse
// @Failure default {object} types.GetImageHashByProductIdResponse
// @Router /api/product-image/get-images-by-product-id [get]
func (h *Handler) getImagesByProductId(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product-image.getImageHashByProductId"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		productIdStr := r.URL.Query().Get("product_id")
		if productIdStr == "" {
			log.Error("missing product_id", slog.String("error", "missing product_id"))
			render.JSON(w, r, response.Error("Missing product_id"))
			return
		}

		productId, err := strconv.ParseInt(productIdStr, 10, 64)
		if err != nil {
			log.Error("invalid product_id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid product_id"))
			return
		}

		hashes, err := h.services.ProductsImages.GetImageHashesByProductId(productId)
		if err != nil {
			log.Error("failed to get image hash", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		//for i := 0; i < len(hashes); i++ {
		//	http.ServeFile(w, r, "uploads/"+hashes[i])
		//}
		render.JSON(w, r, types.GetImageHashByProductIdResponse{
			Hashes: hashes,
		})
	}
}

// @Summary createProductImage
// @Tags Product Image
// @Description creating database record
// @ID create-product-image
// @Accept multipart/form-data
// @Produce json
// @Param product_id formData int64 true "ID продукта"
// @Param image formData file true "Файл изображения"
// @Success 200 {object} types.CreateProductImageResponse
// @Failure 400,404 {object} types.CreateProductImageResponse
// @Failure 500 {object} types.CreateProductImageResponse
// @Failure default {object} types.CreateProductImageResponse
// @Router /api/product-image/create-product-image [post]
func (h *Handler) createProductImage(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product-image.createProductImage"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		// Парсинг входных данных
		err := r.ParseMultipartForm(10 << 20) // 10MB max file size
		if err != nil {
			log.Error("failed to parse multipart form", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid form data"))
			return
		}

		// Получаем и конвертируем product_id в int64
		productIdStr := r.FormValue("product_id")
		if productIdStr == "" {
			render.JSON(w, r, response.Error("Missing product_id"))
			return
		}

		productId, err := strconv.ParseInt(productIdStr, 10, 64)
		if err != nil {
			render.JSON(w, r, response.Error("Invalid product_id"))
			return
		}

		// Получаем файл
		file, header, err := r.FormFile("image")
		if err != nil {
			log.Error("failed to get file", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Failed to retrieve image file"))
			return
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				log.Error("failed to close file", slog.String("error", err.Error()))
			}
		}(file)

		// Читаем содержимое файла для хеширования
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			log.Error("failed to read file", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Failed to read image file"))
			return
		}

		// Генерируем SHA256 хеш файла
		hash := sha256.Sum256(fileBytes)
		hashString := hex.EncodeToString(hash[:])

		// Генерируем путь для сохранения файла
		fileExt := filepath.Ext(header.Filename)
		fileName := hashString + fileExt
		filePath := filepath.Join("uploads", fileName)

		// Создаём директорию, если её нет
		err = os.MkdirAll("uploads", os.ModePerm)
		if err != nil {
			log.Error("failed to create upload directory", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Failed to create upload directory"))
			return
		}

		// Сохраняем файл на сервере
		err = os.WriteFile(filePath, fileBytes, 0644)
		if err != nil {
			log.Error("failed to save file", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Failed to save image file"))
			return
		}

		// Записываем product_id и хеш изображения в БД
		recordId, err := h.services.ProductsImages.CreateProductImage(productId, fileName)
		if err != nil {
			log.Error("failed to create record", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		// Возвращаем ответ
		render.JSON(w, r, types.CreateProductImageResponse{
			RecordId: recordId,
		})
	}
}

// @Summary updateProductImage
// @Tags Product Image
// @Description updating database record
// @ID update-product-image
// @Accept  multipart/form-data
// @Produce  json
// @Param old_hash_name formData string true "Название старого изображения"
// @Param image formData file true "Файл изображения"
// @Success 200 {object} types.CreateProductImageResponse
// @Failure 400,404 {object} types.CreateProductImageResponse
// @Failure 500 {object} types.CreateProductImageResponse
// @Failure default {object} types.CreateProductImageResponse
// @Router /api/product-image/update-product-image [put]
func (h *Handler) updateProductImage(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product-image.updateProductImage"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		//var req types.UpdateProductImageRequest
		//if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		//	log.Error("failed to decode request body", slog.String("error", err.Error()))
		//	render.JSON(w, r, response.Error("Invalid request body"))
		//	return
		//}
		var oldHashName = r.FormValue("old_hash_name")
		if oldHashName == "" {
			log.Error("missing old_hash_name", slog.String("error", "missing old_hash_name"))
			render.JSON(w, r, response.Error("missing old_hash_name"))
			return
		}

		//if err := h.services.ProductsImages.DeleteProductImageByName(oldHashName); err != nil {
		//	log.Error("can't delete from ProductImage table by old_hash_name", slog.String("error", "can't delete from ProductImage table by old_hash_name"))
		//	render.JSON(w, r, response.Error("can't delete from ProductImage table by old_hash_name"))
		//	return
		//}

		// Получаем файл
		file, header, err := r.FormFile("image")
		if err != nil {
			log.Error("failed to get file", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Failed to retrieve image file"))
			return
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				log.Error("failed to close file", slog.String("error", err.Error()))
			}
		}(file)

		// Читаем содержимое файла для хеширования
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			log.Error("failed to read file", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Failed to read image file"))
			return
		}

		// Генерируем SHA256 хеш файла
		hash := sha256.Sum256(fileBytes)
		hashString := hex.EncodeToString(hash[:])

		// Генерируем путь для сохранения файла
		fileExt := filepath.Ext(header.Filename)
		fileName := hashString + fileExt
		filePath := filepath.Join("uploads", fileName)

		// Создаём директорию, если её нет
		err = os.MkdirAll("uploads", os.ModePerm)
		if err != nil {
			log.Error("failed to create upload directory", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Failed to create upload directory"))
			return
		}

		// Сохраняем файл на сервере
		err = os.WriteFile(filePath, fileBytes, 0644)
		if err != nil {
			log.Error("failed to save file", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Failed to save image file"))
			return
		}

		err = h.services.ProductsImages.UpdateProductImage(oldHashName, fileName)
		if err != nil {
			log.Error("failed update record", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.UpdateProductImageResponse{
			Status: "OK",
		})
	}
}

// @Summary deleteProductImageByName
// @Tags Product Image
// @Description deleting database record by name
// @ID delete-product-image-by-name
// @Accept  json
// @Produce  json
// @Param input body types.DeleteProductImageByNameRequest true "Удаляет запись об изображении в базе данных по хэшу имени"
// @Success 200 {object} types.DeleteProductImageByNameResponse
// @Failure 400,404 {object} types.DeleteProductImageByNameResponse
// @Failure 500 {object} types.DeleteProductImageByNameResponse
// @Failure default {object} types.DeleteProductImageByNameResponse
// @Router /api/product-image/delete-product-image-by-name [delete]
func (h *Handler) deleteProductImageByName(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product-image.deleteProductImageByName"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.DeleteProductImageByNameRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		err := os.Remove("uploads/" + req.OldHashName)
		if err != nil {
			log.Error("failed delete file", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		err = h.services.ProductsImages.DeleteProductImageByName(req.OldHashName)
		if err != nil {
			log.Error("failed delete record by hash name", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.DeleteProductImageByNameResponse{
			Status: "OK",
		})
	}
}

// @Summary deleteProductImageById
// @Tags Product Image
// @Description deleting database record by id
// @ID delete-product-image-by-id
// @Accept  json
// @Produce  json
// @Param input body types.DeleteProductImageByIdRequest true "Удаляет запись об изображении в базе данных по id"
// @Success 200 {object} types.DeleteProductImageByIdResponse
// @Failure 400,404 {object} types.DeleteProductImageByIdResponse
// @Failure 500 {object} types.DeleteProductImageByIdResponse
// @Failure default {object} types.DeleteProductImageByIdResponse
// @Router /api/product-image/delete-product-image-by-id [delete]
func (h *Handler) deleteProductImageById(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "routes.product-image.deleteProductImageById"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req types.DeleteProductImageByIdRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Invalid request body"))
			return
		}

		imageName, err := h.services.ProductsImages.GetImageHashByImageId(req.ImageId)
		if err != nil {
			log.Error("failed find image name by image id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		err = os.Remove("uploads/" + imageName)
		if err != nil {
			log.Error("failed delete file", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		err = h.services.ProductsImages.DeleteProductImageById(req.ImageId)
		if err != nil {
			log.Error("failed delete record by record id", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("Internal server error"))
			return
		}

		render.JSON(w, r, types.DeleteProductImageByIdResponse{
			Status: "OK",
		})
	}
}
