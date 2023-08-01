import React, { useState } from 'react';
import styles from '../assets/components/LinkInput.module.css';

const LinkInput = ({
  label,
  placeholder,
  buttonText,
  buttonAction,
  onChangeAction,
  loading,
}) => {
  const [url, setUrl] = useState('');

  function onChangeInput(url) {
    onChangeAction();
    setUrl(url);
  }

  function buttonDisabled() {
    return loading || !url || url === '' || !validateWebsiteUrl();
  }

  function validateWebsiteUrl() {
    const pattern = new RegExp(
      '^([a-zA-Z]+:\\/\\/)?' + // protocol
        '((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|' + // domain name
        '((\\d{1,3}\\.){3}\\d{1,3}))' + // OR IP (v4) address
        '(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*' + // port and path
        '(\\?[;&a-z\\d%_.~+=-]*)?' + // query string
        '(\\#[-a-z\\d_]*)?$', // fragment locator
      'i',
    );
    return pattern.test(url);
  }

  return (
    <div className={styles.container}>
      <span>{label}</span>
      <br />
      <input
        className={styles.input}
        value={url}
        onChange={(e) => onChangeInput(e.target.value)}
        placeholder={placeholder}
      />
      <button
        className={styles.successButton}
        onClick={() => buttonAction(url)}
        disabled={buttonDisabled()}
      >
        {buttonText}
      </button>
    </div>
  );
};

export default LinkInput;
