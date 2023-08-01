import React, { useState } from 'react';
import styles from '../assets/components/LinkInput.module.css';

const LinkInput = ({label, placeholder, buttonText, buttonAction, onChangeAction, loading}) => {
  const [url, setUrl] = useState('')

  function onChangeInput(url) {
    onChangeAction()
    setUrl(url)
  }

  function buttonDisabled() {
    return loading || !url || url === ''
  }

  return (
    <div className={styles.container}>
      <span>{label}</span>
      <br/>
      <input className={styles.input} value={url} onChange={e => onChangeInput(e.target.value)} placeholder={placeholder} />
      <button className={styles.successButton} onClick={() => buttonAction(url)} disabled={buttonDisabled()}>{buttonText}</button>
    </div>
  );
};

export default LinkInput;
