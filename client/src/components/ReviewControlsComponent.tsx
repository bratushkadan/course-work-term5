import React, { useEffect, useState } from 'react';
import type { Order, Review } from '../api/types';
import { useAuth, useReviews } from '../stores';
import { useShallow } from 'zustand/react/shallow';
import { BlockBtn, BlockInput, BlockSmallerText, Btn } from './generic';
import dayjs from 'dayjs';
import ReactModal from 'react-modal';
import { api } from '../api';
import { alertError } from '../util/error';

export const ReviewControlsComponent: React.FC<Order['positions'][number]> = (props) => {
  const token = useAuth((state) => state.token);

  const [review, setReview] = useState<Review | undefined>();
  const { reviews, setReviews } = useReviews(useShallow(({ reviews, setReviews }) => ({ reviews, setReviews })));
  const [showAddModal, setShowAddModal] = useState(false);
  const [showEditModal, setShowEditModal] = useState(false);
  const [addModalRating, setAddModalRating] = useState<number>(7);
  const [editModalRating, setEditModalRating] = useState<number>(0);
  const [addModalText, setAddModalText] = useState<string>('');
  const [editModalText, setEditModalText] = useState<string>('');

  useEffect(() => {
    setReview(reviews.find((review) => review.product_id === props.product_id));
  }, [reviews]);

  useEffect(() => {
    if (review) {
      setEditModalRating(review.rating);
      setEditModalText(review.review_text);
    }
  }, [review]);

  const handleAddReview = () => {
    if (!token || !addModalRating) {
      return;
    }

    api
      .addReview(token, {
        product_id: props.product_id,
        rating: addModalRating,
        review_text: addModalText,
      })
      .then((addedReview) => {
        setReviews([...reviews, addedReview]);
        setAddModalRating(7);
        setAddModalText('');
        setShowAddModal(false);
      })
      .catch(alertError);
  };

  const handleEditReview = () => {
    if (!token || !editModalRating) {
      return;
    }

    api
      .editReview(token, {
        product_id: props.product_id,
        rating: editModalRating,
        review_text: editModalText,
      })
      .then((editedReview) => {
        setReviews([...reviews.filter((review) => review.id !== editedReview.id), editedReview]);
        setShowEditModal(false);
      })
      .catch(alertError);
  };

  const handleDeleteReview = () => {
    if (!token) {
      return;
    }

    api
      .deleteReview(token, {product_id: props.product_id})
      .then((editedReview) => {
        setReviews(reviews.filter(r => r.id !== editedReview.id));
        setShowEditModal(false);
      })
      .catch(alertError);
  }

  return (
    <>
      {review ? (
        <>
          <div>
            <b>Ваш отзыв:</b>
            <br />
            {review.review_text}
            <br />
            <BlockSmallerText>Оценка: {review.rating}/10</BlockSmallerText>
            <BlockSmallerText>
              {dayjs(review.created).format('DD.MM.YYYY HH:mm:ss')}
              {review.created === review.modified
                ? ''
                : ` (изменено ${dayjs(review.modified).format('DD.MM.YYYY HH:mm:ss')})`}
            </BlockSmallerText>
          </div>
          <Btn onClick={() => setShowEditModal(true)}>Редактировать отзыв</Btn>
          <Btn onClick={handleDeleteReview}>Удалить отзыв</Btn>
          <ReactModal
            isOpen={showEditModal}
            shouldFocusAfterRender={true}
            shouldCloseOnOverlayClick={true}
            shouldCloseOnEsc={true}
            ariaHideApp={false}
            onRequestClose={() => setShowEditModal(false)}
            contentLabel="modal"
          >
            <label htmlFor="edit_rating">Оценка:</label>
            <BlockInput
              name="rating"
              id="edit_rating"
              type="number"
              placeholder="оценка (1-10)"
              value={editModalRating}
              onChange={(e) => setEditModalRating(Number(e.currentTarget.value))}
            />
            <label htmlFor="edit_descr">Отзыв:</label>
            <br />
            <textarea
              name="edit_descr"
              id="edit_descr"
              cols={30}
              rows={10}
              value={editModalText}
              onChange={(e) => setEditModalText(e.currentTarget.value)}
            />
            <BlockBtn
              type="submit"
              disabled={editModalText.trim() === '' || !editModalRating}
              onClick={handleEditReview}
            >
              Отправить
            </BlockBtn>
            <BlockBtn
              style={{
                position: 'absolute',
                top: '1rem',
                right: '1rem',
              }}
              onClick={() => setShowEditModal(false)}
            >
              Закрыть окно
            </BlockBtn>
          </ReactModal>
        </>
      ) : (
        <>
          <Btn onClick={() => setShowAddModal(true)}>Оставить отзыв</Btn>
          <ReactModal
            isOpen={showAddModal}
            shouldFocusAfterRender={true}
            shouldCloseOnOverlayClick={true}
            shouldCloseOnEsc={true}
            ariaHideApp={false}
            onRequestClose={() => setShowAddModal(false)}
            contentLabel="modal"
          >
            <label htmlFor="add_rating">Добавить оценку:</label>
            <BlockInput
              name="rating"
              id="add_rating"
              type="number"
              placeholder="оценка (1-10)"
              value={addModalRating}
              onChange={(e) => setAddModalRating(Number(e.currentTarget.value))}
            />
            <label htmlFor="add_descr">Добавить отзыв:</label>
            <br />
            <textarea
              name="add_descr"
              id="add_descr"
              cols={30}
              rows={10}
              value={addModalText}
              onChange={(e) => setAddModalText(e.currentTarget.value)}
            />
            <BlockBtn type="submit" disabled={addModalText.trim() === '' || !addModalRating} onClick={handleAddReview}>
              Отправить
            </BlockBtn>
            <BlockBtn
              style={{
                position: 'absolute',
                top: '1rem',
                right: '1rem',
              }}
              onClick={() => setShowAddModal(false)}
            >
              Закрыть окно
            </BlockBtn>
          </ReactModal>
        </>
      )}
    </>
  );
};
