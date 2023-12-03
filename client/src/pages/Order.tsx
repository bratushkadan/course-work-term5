import { useEffect, useState } from 'react';
import { api } from '../api';
import { OrderInfoComponent } from '../components/OrderInfo';
import { useAuth, useMe, useReviews } from '../stores';
import { Order } from '../api/types';
import { alertError } from '../util/error';
import { useParams } from 'react-router-dom';
import { OrderProductComponent } from '../components/OrderProduct';

export const OrderPage: React.FC = () => {
  const token = useAuth((state) => state.token);
  const me = useMe(state => state.me)
  const [order, setOrder] = useState<Order | null>(null);
  const setReviews = useReviews((state) => state.setReviews);

  const params = useParams();

  useEffect(() => {
    if (token === null) {
      return;
    }

    api
      .getOrder(token, params.id as unknown as number)
      .then(setOrder)
      .catch(alertError);
  }, [token]);

  useEffect(() => {
    if (!me) {
      return
    }

    api.getReviews({ user_id: me.id }).then(setReviews).catch(alertError);
  }, [me]);

  return (
    order && (
      <>
        <OrderInfoComponent {...order} />
        <div>Сумма заказа: {order.positions.reduce((prev, cur) => prev + cur.price, 0) / 100} ₽</div>
        {order.positions.map((props) => (
          <OrderProductComponent key={props.product_id} {...props} price={props.price / 100} isWithProductLink={true} />
        ))}
      </>
    )
  );
};
