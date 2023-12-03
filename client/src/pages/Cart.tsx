import { useShallow } from 'zustand/react/shallow';
import { useAuth, useCart } from '../stores';
import { api } from '../api';
import { useEffect } from 'react';
import { CartPositionComponent } from '../components/CartPosition';

export const CartPage: React.FC = () => {
  const token = useAuth((state) => state.token);

  const { cart, setCart } = useCart(useShallow(({ cart, setCart }) => ({ cart, setCart })));

  useEffect(() => {
    if (token) {
      api.getCart(token).then(setCart);
    }
  }, [token]);

  return (
    <>
      <h1>Корзина{cart.length === 0 ? <span>{' '} пуста</span> : <span></span>}</h1>
      {cart.length > 0 && cart.map((props) => (
        <CartPositionComponent key={props.product_id} {...props} />
      ))}
    </>
  );
};
