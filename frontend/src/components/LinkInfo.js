import React from 'react';
import styles from '../assets/components/LinkInfo.module.css';

const LinkInfo = ({label, link}) => {
  // Don't render if url or shortUrl is empty
  if(link.url === '' || link.shortUrl === '') {
    return
  }

  return (
    <div className={styles.container}>
      <span>{label}</span>
      <br/>
      <label>
        URL corta: {link.shortUrl}
      </label>
      <label>
        URL completa: {link.url}
      </label>
      <label>
        Fecha de creación: {link.createdAt}
      </label>
    </div>
  );
};

export default LinkInfo;
