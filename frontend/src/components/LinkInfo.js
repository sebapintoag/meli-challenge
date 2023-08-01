import React from 'react';
import { Link } from 'react-router-dom';
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
          <label className={styles.label}>
            URL corta: <Link to={link.shortUrl} target="_blank">{link.shortUrl}</Link>
          </label>
          <label className={styles.label}>
            URL completa: <Link to={link.url} target="_blank">{link.url}</Link>
          </label>
          <label className={styles.label}>
            Fecha de creación: {link.createdAt}
          </label>
        </>
      )}
    </div>
  );
};

export default LinkInfo;
