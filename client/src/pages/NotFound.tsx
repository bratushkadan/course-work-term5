import {Link} from 'react-router-dom'
import styled from 'styled-components'

const NotFoundWrapper = styled.div`
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  font-size: 3rem;
  font-weight: bold;
`

const TextLink = styled.div`
  font-size: 2rem;
  color: var(--color-link);
`

export const NotFound: React.FC = () => {
  return <NotFoundWrapper>
      <div>Страница не найдена :{'('}</div>
      <Link to="/"><TextLink>На главную</TextLink></Link>
  </NotFoundWrapper>
}