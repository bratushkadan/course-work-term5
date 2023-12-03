import dayjs from 'dayjs';

import { OrderInfo } from '../api/types';
import { localizeOrderStatus } from '../api/util';
import { Link } from 'react-router-dom';

export const OrderInfoComponent: React.FC<OrderInfo & { isWithOrderLink?: boolean }> = (props) => {
  return (
    <>
      <h2>
        <>
          {props.isWithOrderLink ? <Link to={`/orders/${props.id}`}>Заказ №{props.id}</Link> : <>Заказ №{props.id}</>}
        </>{' '}
        от {dayjs(props.created).format('DD.MM.YYYY HH:mm:ss')}
      </h2>
      <div>
        Статус заказа: {localizeOrderStatus(props.status)}
        {props.created === props.status_modified
          ? ''
          : ` (обновлено ${dayjs(props.status_modified).format('DD.MM.YYYY')} в ${dayjs(props.status_modified).format(
              'HH:mm:ss'
            )})`}
      </div>
    </>
  );
};
