import { useEffect, useState } from 'react';
import { ProductComponent } from '../components/product/Product';
import { useParams } from 'react-router-dom';
import { api } from '../api';
import { Product, Review } from '../api/types';
import { alertError } from '../util/error';
import { ReviewComponent } from '../components/reviews/Review';

export const ProductPage: React.FC = () => {
  const [product, setProduct] = useState<Product | null>(null);
  const [productReviews, setProductReviews] = useState<Review[]>([]);

  const params = useParams();

  useEffect(() => {
    api
      .getProduct(params.id as unknown as number)
      .then(setProduct)
      .catch(alertError);
  }, []);

  useEffect(() => {
    api.getReviews({ product_id: params.id as unknown as number }).then(setProductReviews);
  }, []);

  return (
    <>
      {!product ? null : (
        <>
          <ProductComponent {...product} />
          {productReviews.length === 0 ? (
            <h3>Отзывов на товар нет</h3>
          ) : (
            <>
              <h3>Отзывы на товар</h3>
              {productReviews.map((props) => (
                <ReviewComponent {...props} key={props.id} isDisplayProductName={false} />
              ))}
            </>
          )}
        </>
      )}
    </>
  );
};
