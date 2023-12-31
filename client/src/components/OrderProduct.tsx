import React, { useEffect } from 'react';
import { Link } from 'react-router-dom';
import type { Order, Review } from '../api/types';
import { useMe } from '../stores';
import {ReviewControlsComponent} from './ReviewControlsComponent';

export const OrderProductComponent: React.FC<Order['positions'][number] & { isWithProductLink?: boolean }> = (
  props
) => {
  return (
    <div className="order-product-card">
      <h2>{props.isWithProductLink ? <Link to={`/products/${props.product_id}`}>{props.name}</Link> : props.name}</h2>
      <img width={150} height={150} src={props.image_url} alt={props.name} />
      <div>{props.description}</div>
      <div>{props.price} ₽</div>
      <div>
        Категория: <Link to={`/?category_id=${props.category_id}`}>{props.category_name}</Link>
      </div>
      {/* !!! */}
      <div>
        Продавец: <Link to={`/stores/${props.store_id}`}>{props.store_name}</Link>
      </div>
      <hr />
      <ReviewControlsComponent {...props} />
    </div>
  );
};
