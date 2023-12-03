import dayjs from 'dayjs';
import { Link } from 'react-router-dom';

import type { Review as IReview } from '../../api/types';
import styled from 'styled-components';
import {BlockSmallerText} from '../generic';

const ReviewWrapper = styled.div`
  border: 1px solid black;
  padding: 5px;
  max-width: 1000px;
`;

export const ReviewComponent: React.FC<IReview & { isDisplayProductName?: boolean }> = (props) => {
  return (
    <ReviewWrapper>
      <div>
        {props.user_name} <BlockSmallerText>{dayjs(props.created).format('DD.MM.YYYY HH:mm:ss')}{' '}
        {props.modified === props.created ? '' : `(изменено ${dayjs(props.modified).format('DD.MM.YYYY HH:mm:ss')})`}</BlockSmallerText>
      </div>
      {props.isDisplayProductName ?? true ? (
        <div>
          <Link to={`/products/${props.product_id}`}>{props.product_name}</Link>
        </div>
      ) : null}
      <BlockSmallerText>Оценка: {props.rating}/10</BlockSmallerText>
      <div>{props.review_text}</div>
    </ReviewWrapper>
  );
};
