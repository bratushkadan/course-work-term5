import ky from 'ky';

import type {
  User,
  CreateUserPayload,
  GetTokenPayload,
  Token,
  ValidateTokenPayload,
  ValidateTokenResponse,
  CreateUserResponse,
  GetProductsPayload,
  Product,
  Store,
  Category,
  ChangeCartPositionPayload,
  MinimalCartPosition,
  Order,
  OrderInfo,
  AddDeleteFavoriteResponse,
  Favorite,
  GetReviewsPayload,
  Review,
  AddReviewPayload,
  DeleteReviewPayload,
  EditReviewPayload,
  CartPosition,
} from './types';
import { omitFalsy, transformProduct } from './util';

const v1Api = ky.extend({
  prefixUrl: 'http://localhost:8080/v1',
});

const X_AUTH_TOKEN = 'X-Auth-Token';

const getUser = (id: User['id']) => v1Api.get(`users/${id}`).json<CreateUserResponse>();
const getUserMeByToken = (token: string) =>
  v1Api
    .get(`users/me`, {
      headers: {
        [X_AUTH_TOKEN]: token,
      },
    })
    .json<CreateUserResponse>();
const createUser = (payload: CreateUserPayload) =>
  v1Api
    .post(`users`, {
      json: payload,
    })
    .json<User>();
const getToken = (payload: GetTokenPayload) =>
  v1Api.get(`auth/token/user?` + new URLSearchParams(payload)).json<Token>();
const validateToken = (payload: ValidateTokenPayload) =>
  v1Api
    .post(`auth/token/user`, {
      headers: {
        [X_AUTH_TOKEN]: payload.token,
      },
    })
    .json<ValidateTokenResponse>();

const getProducts = async ({ filter = {}, sort = {} }: GetProductsPayload = {}) => {
  const queryParams: Record<string, string> = {};
  for (const [key, val] of Object.entries(sort)) {
    queryParams[`sort_${key}`] = val;
  }
  for (const [key, val] of Object.entries(filter)) {
    queryParams[`filter.${key}`] = String(val);
  }
  const urlSearhParams = new URLSearchParams(queryParams);

  const products = await v1Api.get(`products?${urlSearhParams}`).json<Product[]>();

  return products.map(transformProduct);
};
const getProduct = (id: number) => v1Api.get(`products/${id}`).json<Product>().then(transformProduct);

const getStores = () => v1Api.get('stores').json<Store[]>();
const getStore = (id: number) => v1Api.get(`stores/${id}`).json<Store>();

const getCategories = () => v1Api.get('categories').json<Category[]>();
const getCategory = (id: number) => v1Api.get(`categories/${id}`).json<Category>();

export const api = {
  getUser,
  getUserMeByToken,
  createUser,
  getToken,
  validateToken,
  getProducts,
  getProduct,
  getStores,
  getStore,
  getCategories,
  getCategory,
  getCart: (token: string) =>
    v1Api
      .get('cart', {
        headers: {
          [X_AUTH_TOKEN]: token,
        },
      })
      .json<CartPosition[]>(),
  changeCartPosition: (token: string, payload: ChangeCartPositionPayload) =>
    v1Api
      .post('cart', {
        headers: {
          [X_AUTH_TOKEN]: token,
        },
        json: payload,
      })
      .json<MinimalCartPosition>(),
  clearCart: (token: string) =>
    v1Api
      .delete('cart', {
        headers: {
          [X_AUTH_TOKEN]: token,
        },
      })
      .json<MinimalCartPosition[]>(),
  getOrder: (token: string, id: number) =>
    v1Api
      .get(`orders/${id}`, {
        headers: {
          [X_AUTH_TOKEN]: token,
        },
      })
      .json<Order>(),
  getOrders: (token: string) =>
    v1Api
      .get(`orders`, {
        headers: {
          [X_AUTH_TOKEN]: token,
        },
      })
      .json<OrderInfo[]>(),
  createOrder: (token: string) =>
    v1Api
      .post('orders', {
        headers: {
          [X_AUTH_TOKEN]: token,
        },
      })
      .json<Order>(),
  getFavorites: (token: string) =>
    v1Api
      .get('favorites', {
        headers: {
          [X_AUTH_TOKEN]: token,
        },
      })
      .json<Favorite[]>(),
  addFavorite: (token: string, productId: number) =>
    v1Api
      .post(`favorites?product_id=${productId}`, {
        headers: {
          [X_AUTH_TOKEN]: token,
        },
      })
      .json<AddDeleteFavoriteResponse>(),
  deleteFavorite: (token: string, productId: number) =>
    v1Api
      .delete(`favorites?product_id=${productId}`, {
        headers: {
          [X_AUTH_TOKEN]: token,
        },
      })
      .json<AddDeleteFavoriteResponse>(),
  getReviews: (payload: GetReviewsPayload) =>
    v1Api.get(`reviews?${new URLSearchParams(payload as unknown as URLSearchParams)}`).json<Review[]>(),
  addReview: (token: string, payload: AddReviewPayload) =>
    v1Api
      .post(`reviews`, {
        headers: {
          [X_AUTH_TOKEN]: token,
        },
        json: payload,
      })
      .json<Review>(),
  editReview: (token: string, payload: EditReviewPayload) =>
    v1Api
      .patch(`reviews`, {
        headers: {
          [X_AUTH_TOKEN]: token,
        },
        json: payload,
      })
      .json<Review>(),
  deleteReview: (token: string, payload: DeleteReviewPayload) =>
    v1Api
      .delete(`reviews?product_id=${payload.product_id}`, {
        headers: {
          [X_AUTH_TOKEN]: token,
        },
      })
      .json<{id: number}>(),
};
