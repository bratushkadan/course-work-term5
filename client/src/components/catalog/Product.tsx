import React from 'react';
import { Link } from 'react-router-dom';
import type { Product as IProduct } from '../../api/types';

export const Product: React.FC<IProduct> = (props) => {
  return (
    <div className="product-card">
      <h2>{props.name}</h2>
      <img width={200} height={200} src={props.image_url} alt={props.name} />
      <p>{props.description}</p>
      <p>{props.price} ₽</p>
      <p>Категория: <Link to={`/?filter.category_id=${props.category.id}`}>{props.category.name}</Link></p>
      {/* !!! */}
      <p>Продавец: <Link to={`/stores/${props.store_id}`}>{props.store_name}</Link></p>
    </div>
  );
};
