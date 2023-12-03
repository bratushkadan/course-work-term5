import dayjs from 'dayjs';
import { Link } from 'react-router-dom';

import type { Store as IStore } from '../../api/types';

export const Store: React.FC<IStore & { isWithLinkToStore?: boolean }> = (props) => {
  return (
    <>
      <h1>{props.isWithLinkToStore ? <Link to={`/stores/${props.id}`}>{props.name}</Link> : props.name}</h1>
      <p>{props.email}</p>
      <p>На Floral с {dayjs(props.created).format('DD.MM.YYYY')}</p>
      <p>
        <Link to={`/?filter.store_id=${props.id}`}>Продукция магазина</Link>
      </p>
    </>
  );
};
