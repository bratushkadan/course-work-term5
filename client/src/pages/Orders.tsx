import { useEffect, useState } from 'react';
import { api } from '../api';
import { OrderInfoComponent } from '../components/OrderInfo';
import { useAuth } from '../stores';
import { OrderInfo } from '../api/types';
import { Link } from 'react-router-dom';

export const OrdersPage: React.FC = () => {
  const token = useAuth((state) => state.token);
  const [orders, setOrders] = useState<OrderInfo[]>([]);

  useEffect(() => {
    if (token) {
      api.getOrders(token).then(setOrders);
    }
  }, [token]);

  return (
    <>
      <h1>
        {orders.length === 0 ? (
          <>
            Заказов еще нет, <Link to="/">сделаем?</Link>
          </>
        ) : (
          'Заказы'
        )}
      </h1>
      {orders.map((props) => (
        <OrderInfoComponent key={props.id} {...props} isWithOrderLink={true} />
      ))}
    </>
  );
};
