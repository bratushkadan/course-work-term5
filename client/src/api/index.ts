import ky from 'ky';

import type {
  User,
  CreateUserPayload,
  GetTokenPayload,
  Token,
  ValidateTokenPayload,
  ValidateTokenResponse,
  CreateUserResponse,
} from './types';

const v1Api = ky.extend({
  prefixUrl: 'http://localhost:8080/v1',
});

const X_AUTH_TOKEN = 'X-Auth-Token';

const getUser = (id: User['id']) => v1Api.get(`users/${id}`).json<CreateUserResponse>();
const getUserMeByToken = (token: string) => v1Api.get(`users/me`, {
  headers: {
    [X_AUTH_TOKEN]: token,
  }
}).json<CreateUserResponse>();
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

export const api = {
  getUser,
  getUserMeByToken,
  createUser,
  getToken,
  validateToken,
};
