import React, {useCallback} from 'react';
import {CartPosition} from '../api/types';
import {BlockBtn, Btn} from './generic';
import {useShallow} from 'zustand/react/shallow';
import {useAuth, useCart, useMe} from '../stores';
import {api} from '../api';
import {alertError} from '../util/error';


export const CartControls: React.FC<CartPosition> = (props) => {
  const token = useAuth((state) => state.token);
  const {cart, setCart} = useCart(useShallow((state) => ({cart: state.cart, setCart: state.setCart})));

  if (!token) {
    return null;
  }

  const handleAddToCart = useCallback(() => {
    api 
      .changeCartPosition(token, {
        product_id: props.product_id,
        quantity: 1,
      })
      .then((updatedCartPosition) => setCart([
        ...cart,
        {
          product_id: props.product_id,
          name: props.name,
          description: props.description,
          quantity: updatedCartPosition.quantity,
          category_id: props.category_id,
          category_name: props.category_name,
          store_id: props.store_id,
          store_name: props.store_name,
          price: props.price,
          image_url: props.image_url,
        },
      ])
      );
  }, [cart, props.product_id]);

  const handleIncrementCartPosition = useCallback(
    (cartPosition: CartPosition) => {
      api
        .changeCartPosition(token, {
          product_id: props.product_id,
          quantity: cartPosition.quantity + 1,
        })
        .then((updatedCartPosition) => setCart(
          cart.map((cartPosition) => {
            if (cartPosition.product_id !== updatedCartPosition.product_id) {
              return cartPosition;
            }
            return {...cartPosition, quantity: updatedCartPosition.quantity};
          })
        )
        )
        .catch(alertError);
    },
    [cart, props.product_id]
  );
  const handleDecrementCartPosition = useCallback(
    (cartPosition: CartPosition) => {
      api
        .changeCartPosition(token, {
          product_id: props.product_id,
          quantity: cartPosition.quantity - 1,
        })
        .then((updatedCartPosition) => setCart(
          cart
            .map((cartPosition) => {
              if (cartPosition.product_id !== updatedCartPosition.product_id) {
                return cartPosition;
              }
              return {...cartPosition, quantity: updatedCartPosition.quantity};
            })
            .filter((cartPosition) => cartPosition.quantity !== 0)
        )
        )
        .catch(alertError);
    },
    [cart, props.product_id]
  );

  const item = cart.find((cartPosition) => cartPosition.product_id === props.product_id);

  return (
    <>
      {!item ? (
        <BlockBtn onClick={handleAddToCart}>Добавить в корзину</BlockBtn>
      ) : (
        <div>
          <span style={{display: 'inline-block', marginRight: '0.5rem'}}>В корзине: {item.quantity}</span>
          <Btn onClick={() => handleIncrementCartPosition(item)}>+</Btn>
          <Btn onClick={() => handleDecrementCartPosition(item)}>-</Btn>
        </div>
      )}
    </>
  );
};
