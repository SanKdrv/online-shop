import React, { useState, useEffect } from 'react';
import axios from 'axios';

response.data.category_id = undefined;

const AdminCategory = () => {
    const [categories, setCategories] = useState([]);
    const [newCategory, setNewCategory] = useState('');
    const [editingId, setEditingId] = useState(null);
    const [editValue, setEditValue] = useState('');
    const [loading, setLoading] = useState({
        fetch: false,
        create: false,
        update: false,
        delete: false
    });
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');

    // Настройка axios
    const api = axios.create({
        baseURL: 'http://localhost:8082/api',
        // headers: {
        //     'Content-Type': 'application/json',
        //     'Authorization': `Bearer ${localStorage.getItem('token')}` // Если нужна авторизация
        // }
    });

    // Получение категорий
    const fetchCategories = async () => {
        try {
            setLoading(prev => ({...prev, fetch: true}));
            setError('');

            const response = await api.get('/category/get-all-categories');

            // Обработка разных форматов ответа
            let receivedCategories = [];
            if (Array.isArray(response.data)) {
                receivedCategories = response.data;
            } else if (response.data && typeof response.data === 'object') {
                receivedCategories = response.data.categories || response.data.items || [];
            }

            setCategories(receivedCategories);
        } catch (err) {
            handleApiError(err, 'Не удалось загрузить категории');
        } finally {
            setLoading(prev => ({...prev, fetch: false}));
        }
    };

    // Создание категории
    const handleCreateCategory = async () => {
        if (!newCategory.trim()) {
            setError('Название категории не может быть пустым');
            return;
        }

        try {
            setLoading(prev => ({...prev, create: true}));
            setError('');

            const response = await api.post('/category/create-category', {
                name: newCategory.trim()
            });

            if (response.data && response.data.category_id) {
                setCategories(prev => [...prev, response.data]);
                setNewCategory('');
                showSuccess('Категория успешно создана');
            }
        } catch (err) {
            handleApiError(err, 'Ошибка при создании категории');
        } finally {
            setLoading(prev => ({...prev, create: false}));
        }
    };

    // Обновление категории
    const handleUpdateCategory = async () => {
        if (!editValue.trim()) {
            setError('Название категории не может быть пустым');
            return;
        }

        try {
            setLoading(prev => ({...prev, update: true}));
            setError('');

            const response = await api.put('/category/update-category', {
                id: editingId,
                name: editValue.trim()
            });

            if (response.data) {
                setCategories(prev =>
                    prev.map(cat =>
                        cat.id === editingId ? {...cat, name: editValue.trim()} : cat
                    )
                );
                cancelEditing();
                showSuccess('Категория успешно обновлена');
            }
        } catch (err) {
            handleApiError(err, 'Ошибка при обновлении категории');
        } finally {
            setLoading(prev => ({...prev, update: false}));
        }
    };

    // Удаление категории
    const handleDeleteCategory = async (id) => {
        if (!window.confirm('Вы уверены, что хотите удалить эту категорию?')) return;

        try {
            setLoading(prev => ({...prev, delete: true}));
            setError('');

            await api.delete('/category/delete-category', {
                data: { id }
            });

            setCategories(prev => prev.filter(cat => cat.id !== id));
            showSuccess('Категория успешно удалена');
        } catch (err) {
            handleApiError(err, 'Ошибка при удалении категории');
        } finally {
            setLoading(prev => ({...prev, delete: false}));
        }
    };

    // Вспомогательные функции
    const startEditing = (category) => {
        setEditingId(category.id);
        setEditValue(category.name);
        setError('');
    };

    const cancelEditing = () => {
        setEditingId(null);
        setEditValue('');
    };

    const showSuccess = (message) => {
        setSuccess(message);
        setTimeout(() => setSuccess(''), 3000);
    };

    const handleApiError = (err, defaultMessage) => {
        let errorMessage = defaultMessage;

        if (err.response) {
            // Обработка 405 ошибки
            if (err.response.status === 405) {
                errorMessage = 'Метод не разрешен. Проверьте endpoint.';
            } else {
                errorMessage = err.response.data?.message || err.response.statusText || defaultMessage;
            }
        } else if (err.request) {
            errorMessage = 'Нет ответа от сервера';
        } else {
            errorMessage = err.message || defaultMessage;
        }

        setError(errorMessage);
        console.error('API Error:', err);
    };

    useEffect(() => {
        fetchCategories();
    }, []);

    return (
        <div style={styles.container}>
            <h1 style={styles.title}>Управление категориями</h1>

            {/* Сообщения об ошибках и успехе */}
            {error && <div style={styles.error}>{error}</div>}
            {success && <div style={styles.success}>{success}</div>}

            {/* Форма добавления */}
            <div style={styles.formContainer}>
                <input
                    type="text"
                    value={newCategory}
                    onChange={(e) => setNewCategory(e.target.value)}
                    placeholder="Новая категория"
                    style={styles.input}
                    disabled={loading.create}
                />
                <button
                    onClick={handleCreateCategory}
                    style={styles.addButton}
                    disabled={loading.create}
                >
                    {loading.create ? 'Создание...' : 'Добавить'}
                </button>
            </div>

            {/* Таблица категорий */}
            {loading.fetch ? (
                <div style={styles.loading}>Загрузка категорий...</div>
            ) : (
                <table style={styles.table}>
                    <thead>
                    <tr>
                        <th style={styles.th}>ID</th>
                        <th style={styles.th}>Название</th>
                        <th style={styles.th}>Действия</th>
                    </tr>
                    </thead>
                    <tbody>
                    {categories.length > 0 ? (
                        categories.map(category => (
                            <tr key={category.id} style={styles.tr}>
                                <td style={styles.td}>{category.id}</td>
                                <td style={styles.td}>
                                    {editingId === category.id ? (
                                        <input
                                            type="text"
                                            value={editValue}
                                            onChange={(e) => setEditValue(e.target.value)}
                                            style={styles.editInput}
                                        />
                                    ) : (
                                        category.name
                                    )}
                                </td>
                                <td style={styles.td}>
                                    {editingId === category.id ? (
                                        <>
                                            <button
                                                onClick={handleUpdateCategory}
                                                style={{...styles.button, ...styles.saveButton}}
                                                disabled={loading.update}
                                            >
                                                {loading.update ? 'Сохранение...' : 'Сохранить'}
                                            </button>
                                            <button
                                                onClick={cancelEditing}
                                                style={{...styles.button, ...styles.cancelButton}}
                                            >
                                                Отмена
                                            </button>
                                        </>
                                    ) : (
                                        <>
                                            <button
                                                onClick={() => startEditing(category)}
                                                style={{...styles.button, ...styles.editButton}}
                                            >
                                                Редактировать
                                            </button>
                                            <button
                                                onClick={() => handleDeleteCategory(category.id)}
                                                style={{...styles.button, ...styles.deleteButton}}
                                                disabled={loading.delete}
                                            >
                                                {loading.delete ? 'Удаление...' : 'Удалить'}
                                            </button>
                                        </>
                                    )}
                                </td>
                            </tr>
                        ))
                    ) : (
                        <tr>
                            <td colSpan="3" style={styles.noData}>Нет категорий</td>
                        </tr>
                    )}
                    </tbody>
                </table>
            )}
        </div>
    );
};

// Стили компонента
const styles = {
    container: {
        maxWidth: '1000px',
        margin: '0 auto',
        padding: '20px',
        fontFamily: 'Arial, sans-serif'
    },
    title: {
        textAlign: 'center',
        color: '#333',
        marginBottom: '30px'
    },
    error: {
        color: '#d32f2f',
        backgroundColor: '#fde0e0',
        padding: '10px',
        borderRadius: '4px',
        marginBottom: '20px',
        border: '1px solid #f5c6cb'
    },
    success: {
        color: '#388e3c',
        backgroundColor: '#e8f5e9',
        padding: '10px',
        borderRadius: '4px',
        marginBottom: '20px',
        border: '1px solid #c3e6cb'
    },
    formContainer: {
        display: 'flex',
        marginBottom: '30px'
    },
    input: {
        flex: 1,
        padding: '10px',
        fontSize: '16px',
        border: '1px solid #ddd',
        borderRadius: '4px 0 0 4px',
        outline: 'none'
    },
    addButton: {
        padding: '10px 20px',
        backgroundColor: '#4caf50',
        color: 'white',
        border: 'none',
        borderRadius: '0 4px 4px 0',
        cursor: 'pointer',
        fontSize: '16px',
        transition: 'background-color 0.3s',
        ':hover': {
            backgroundColor: '#388e3c'
        }
    },
    loading: {
        textAlign: 'center',
        padding: '20px',
        color: '#666'
    },
    table: {
        width: '100%',
        borderCollapse: 'collapse',
        boxShadow: '0 1px 3px rgba(0,0,0,0.1)'
    },
    th: {
        backgroundColor: '#f5f5f5',
        padding: '12px 15px',
        textAlign: 'left',
        borderBottom: '1px solid #ddd'
    },
    tr: {
        borderBottom: '1px solid #eee',
        ':hover': {
            backgroundColor: '#f9f9f9'
        }
    },
    td: {
        padding: '12px 15px',
        verticalAlign: 'middle'
    },
    editInput: {
        width: '100%',
        padding: '8px',
        border: '1px solid #ddd',
        borderRadius: '4px'
    },
    button: {
        padding: '6px 12px',
        marginRight: '5px',
        border: 'none',
        borderRadius: '4px',
        cursor: 'pointer',
        fontSize: '14px',
        transition: 'all 0.3s'
    },
    editButton: {
        backgroundColor: '#2196f3',
        color: 'white',
        ':hover': {
            backgroundColor: '#0b7dda'
        }
    },
    deleteButton: {
        backgroundColor: '#f44336',
        color: 'white',
        ':hover': {
            backgroundColor: '#d32f2f'
        }
    },
    saveButton: {
        backgroundColor: '#4caf50',
        color: 'white',
        ':hover': {
            backgroundColor: '#388e3c'
        }
    },
    cancelButton: {
        backgroundColor: '#ff9800',
        color: 'white',
        ':hover': {
            backgroundColor: '#f57c00'
        }
    },
    noData: {
        textAlign: 'center',
        padding: '20px',
        color: '#666'
    }
};

export default AdminCategory;