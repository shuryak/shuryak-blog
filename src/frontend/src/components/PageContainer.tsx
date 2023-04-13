import React, {useState, PropsWithChildren} from 'react'
import { useNavigate } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faChevronLeft, faChevronRight, faExpand, faCompress } from '@fortawesome/free-solid-svg-icons'
import '@styles/page.scss'

export const PageContainer: React.FunctionComponent<PropsWithChildren<any>> = (props: PropsWithChildren<any>) => {
  const [isExpand, setIsExpand] = useState<boolean>(false)
  const navigate = useNavigate();

  return (
    <>
      <div className={'page' + (isExpand ? ' expanded-page' : '')}>
        <div className="page-options-section">
          <p className="page-options-section-text">Здесь должен быть поиск...</p>
          <div className="user-input btn ctrl" onClick={() => navigate(-1)}>
            <FontAwesomeIcon icon={faChevronLeft} />
          </div>
          <div className="user-input btn ctrl" onClick={() => navigate(1)}>
            <FontAwesomeIcon icon={faChevronRight} />
          </div>
          <div className="user-input btn ctrl" id="expand-btn" onClick={() => setIsExpand(prev => !prev)}>
            <FontAwesomeIcon icon={isExpand ? faCompress : faExpand} />
          </div>
        </div>
        <article>
          {props.children}
        </article>
      </div>
    </>
  )
}

export default PageContainer
