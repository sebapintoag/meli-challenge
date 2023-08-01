import React from 'react';
import styles from '../assets/components/LinkInfo.module.css';

const LinkInfo = ({ label, link, visible, showDetails }) => {
  // Don't render if component is not visible
  if (!visible) {
    return;
  }

  return (
    <div className={styles.container}>
      <b>{label}</b>
      {showDetails && (
        <>
          <br />
          <label>URL corta: {link.shortUrl}</label>
          <label>URL completa: {link.url}</label>
          <label>Fecha de creación: {link.createdAt}</label>
        </>
      )}
    </div>
  );
};

export default LinkInfo;
