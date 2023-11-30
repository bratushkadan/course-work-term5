import React, { useEffect, useLayoutEffect, useState } from 'react';
import {RotateLoader} from 'react-spinners'
import { Login } from './modals/Login';
import { SignUp } from './modals/SignUp';
import { useAuth, useLoginModal, useMe, useSignUpModal } from '../stores';
import { eraseCookie, getCookie, setCookie } from '../util/cookie';
import { BlockBtn, Btn } from './generic';
import styled from 'styled-components';
import { useShallow } from 'zustand/react/shallow';
import { api } from '../api';
import {alertError} from '../util/error';

const LoginWrapper = styled.div`
  display: flex;
  justify-content: flex-end;
  align-items: center;
  flex-grow: 1;
`;

const AUTH_TOKEN_COOKIE_NAME = 'auth_token';

export const LoginPanel: React.FC<{}> = () => {
  const [isTokenValidated, setTokenIsValidated] = useState(false);
  const [isMeRequested, setIsMeRequested] = useState(false);

  const setIsOpenLoginModal = useLoginModal((state) => state.setIsOpen);
  const setIsOpenSignUpModal = useSignUpModal((state) => state.setIsOpen);
  const { me, setMe } = useMe(
    useShallow((state) => ({
      me: state.me,
      setMe: state.setMe,
    }))
  );
  const { token, setToken, clearToken } = useAuth(
    useShallow((state) => ({
      token: state.token,
      setToken: state.setToken,
      clearToken: state.clearToken,
    }))
  );

  useLayoutEffect(() => {
    const tokenCookie = getCookie(AUTH_TOKEN_COOKIE_NAME);
    if (tokenCookie !== null) {
      api
        .validateToken({ token: tokenCookie })
        .then(({ valid }) => {
          if (valid) {
            setTokenIsValidated(true);
            setToken(tokenCookie);
          }
        })
        .finally(() => setTokenIsValidated(true));
    } else {
      setTokenIsValidated(true);
    }
  }, []);
  useEffect(() => {
    if (!isTokenValidated) {
      return;
    }
    if (token === null) {
      eraseCookie(AUTH_TOKEN_COOKIE_NAME);
    } else {
      setCookie(AUTH_TOKEN_COOKIE_NAME, token, 30);
    }
  }, [token, isTokenValidated]);

  useEffect(() => {
    if (token !== null) {
      api.getUserMeByToken(token).then((data) => {
        setMe(data)
        setIsMeRequested(true)
      }).catch(err => {
        alertError(err)
        setIsMeRequested(true)
      });
    }
  }, [token]);

  return (
    <LoginWrapper>
      {token && me !== null ? (
        <>
          <div style={{ marginRight: '.5rem' }}>
            Добро пожаловать, {me.first_name} {me.last_name}!
          </div>
          <BlockBtn
            style={{marginRight: '.5rem'}}
            onClick={clearToken}
          >
            Выйти
          </BlockBtn>
        </>
      ) : (
        <div>
          <RotateLoader loading={!isTokenValidated && !isMeRequested}></RotateLoader>
          {isTokenValidated && ((token === null) || (token && isMeRequested)) && (
            <>
              <Btn onClick={() => setIsOpenLoginModal(true)}>Войти в учетную запись</Btn>
              <Btn onClick={() => setIsOpenSignUpModal(true)}>Зарегистрироваться</Btn>
            </>
          )}
        </div>
      )}
      <Login />
      <SignUp />
    </LoginWrapper>
  );
};
