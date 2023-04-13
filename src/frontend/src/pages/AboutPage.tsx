import React from 'react'
import '@styles/page.scss'

export const AboutPage: React.FC = () => {
  return (
    <>
      <h1>👨‍💻 Александр Коновалов</h1>
      <p>
        Привет! Меня зовут Саша, мне 19 лет, я backend-разработчик. Родился в Марий Эл, сейчас учусь в Санкт-Петербурге.
        В этом блоге я публикую статьи про программирование.
      </p>
      <p>
        Мой обычный стек: <mark>Go</mark>, <mark>Docker</mark>, <mark>Kubernetes</mark>, <mark>PostgreSQL</mark>.
        Постоянно изучаю новые технологии.
      </p>
      <p>
        Telegram: <a href="https://t.me/shuryak">@shuryak</a>. Веду YouTube-канал: <a href="">ШУРЯК</a>. GitHub: <a
        href="https://github.com/shuryak">@shuryak</a>
      </p>
    </>
  )
}

export default AboutPage
