import { useEffect, useState } from 'react';
import type { Store as IStore, Review as IReview } from '../api/types';
import { api } from '../api';
import { Store as StoreComponent } from '../components/stores/Store';
import { useParams } from 'react-router-dom';
import { alertError } from '../util/error';
import { ReviewComponent } from '../components/reviews/Review';

export const StorePage: React.FC = () => {
  const [store, setStore] = useState<IStore | null>();
  const [storeReviews, setStoreReviews] = useState<IReview[]>([]);

  const params = useParams();

  useEffect(() => {
    api
      .getStore(params.id as unknown as number)
      .then(setStore)
      .catch(alertError);
  }, [params.id]);

  useEffect(() => {
    api.getReviews({ store_id: params.id as unknown as number }).then(setStoreReviews);
  }, []);

  return (
    <>
      {store && (
        <>
          <StoreComponent {...store} />
          <h3>{storeReviews.length === 0 ? 'Отзывов нет' : 'Отзывы на продукцию магазина'}</h3>
          {storeReviews.map((props) => (
            <ReviewComponent {...props} key={props.id} isDisplayProductName={true} />
          ))}
        </>
      )}
    </>
  );
};
