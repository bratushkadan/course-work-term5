import React from 'react';
import { CartPosition } from '../api/types';
import { Link } from 'react-router-dom';
import { BlockSmallerText } from './generic';
import {CartControls} from './CartControls';

export const CartPositionComponent: React.FC<CartPosition> = (props) => {
  return (
    <div className="cart-position-card">
      <h3>
        <Link to={`/products/${props.product_id}`}>{props.name}</Link>
      </h3>
      <img width={150} height={150} src={props.image_url} alt={props.name} />
      <div>{props.description}</div>
      <div>{props.price/100} ₽</div>
      <div>
        Категория: <Link to={`/?category_id=${props.category_id}`}>{props.category_name}</Link>
      </div>
      {/* !!! */}
      <BlockSmallerText>Продавец: <Link to={`/stores/${props.store_id}`}>{props.store_name}</Link></BlockSmallerText>
      <CartControls {...props} />
    </div>
  );
};
