import React from 'react';
import { Link } from 'react-router-dom';
import type { Favorite } from '../api/types';
import { api } from '../api';
import { useShallow } from 'zustand/react/shallow';
import { useAuth, useFavorite } from '../stores';
import { alertError } from '../util/error';
import { Btn } from './generic';

const FavoriteControls: React.FC<Favorite> = (props) => {
  const token = useAuth((state) => state.token);
  const { favorites, setFavorites } = useFavorite(
    useShallow(({ favorites, setFavorites }) => ({ favorites, setFavorites }))
  );

  if (!token) {
    return null;
  }

  const handleAddFavorite = () => {
    api
      .addFavorite(token, props.product_id)
      .then(() => {
        setFavorites([
          ...favorites,
          {
            product_id: props.product_id,
            store_id: props.store_id,
            store_name: props.store_name,
            description: props.description,
            name: props.name,
            category_id: props.category_id,
            category_name: props.category_name,
            added_favorite: Date.now(),
            image_url: props.image_url,
            price: props.price,
          },
        ]);
      })
      .catch(alertError);
  };

  const handleRemoveFavorite = () => {
    api
      .deleteFavorite(token, props.product_id)
      .then((deletedFavorite) => {
        setFavorites(favorites.filter((fav) => fav.product_id !== deletedFavorite.id));
      })
      .catch(alertError);
  };

  return (
    <>
      {favorites.find((fav) => fav.product_id === props.product_id) ? (
        <Btn onClick={handleRemoveFavorite}>Удалить из избранного</Btn>
      ) : (
        <Btn onClick={handleAddFavorite}>Добавить в избранное</Btn>
      )}
    </>
  );
};

export const FavoriteProductComponent: React.FC<Favorite & { isWithProductLink?: boolean }> = (props) => {
  return (
    <div className="fav-product-card">
      <h2>{props.isWithProductLink ? <Link to={`/products/${props.product_id}`}>{props.name}</Link> : props.name}</h2>
      <img width={150} height={150} src={props.image_url} alt={props.name} />
      <div>{(props.price / 100)} ₽</div>
      <div>
        Категория: <Link to={`/?category_id=${props.category_id}`}>{props.category_name}</Link>
      </div>
      <div>
        Продавец: <Link to={`/stores/${props.store_id}`}>{props.store_name}</Link>
      </div>
      <FavoriteControls {...props} />
    </div>
  );
};
