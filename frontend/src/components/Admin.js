import React from 'react';

const Admin = () => {
    return (
        <div>
            <nav>
                <ul>
                    <li>
                        <a href="/brand">Таблица брендов</a>
                    </li>
                    <li>
                        <a href="/cart-content">Таблица содержимого корзины</a>
                    </li>
                    <li>
                        <a href="/category">Таблица категорий</a>
                    </li>
                    <li>
                        <a href="/order">Таблица заказов</a>
                    </li>
                    <li>
                        <a href="/order-content">Таблица содержимого заказов</a>
                    </li>
                    <li>
                        <a href="/product">Таблица товаров</a>
                    </li>
                    <li>
                        <a href="/product-image">Таблица изображений товаров</a>
                    </li>
                    <li>
                        <a href="/refresh-session">Таблица сессий пользователей</a>
                    </li>
                    <li>
                        <a href="/user">Таблица пользователей</a>
                    </li>
                </ul>
            </nav>
            <div>
                <h1>Админка онлайн магазина</h1>
                <p>Здесь будет админ панель.</p>
            </div>

        </div>
    );

};

export default Admin;