import React, { useState } from 'react';
import LinkInput from '../components/LinkInput';
import LinkInfo from '../components/LinkInfo';
import axios from 'axios';

const Delete = () => {
  const [label, setLabel] = useState('');
  const [loading, setLoading] = useState(false);
  const [showLinkInfo, setshowLinkInfo] = useState(false);

  function deleteShortUrl(short_url) {
    setLoading(true);
    axios
      .delete('http://localhost/api/v1/delete', {
        data: { short_url: short_url }
      })
      .then((res) => {
        setLabel(`URL corta ${short_url} eliminada exitosamente`);
      })
      .catch((err) => {
        setLabel(`No se pudo eliminar la URL corta ${short_url}`);
      })
      .then(() => {
        setLoading(false);
        setshowLinkInfo(true);
      });
  }

  function restartLink() {
		setshowLinkInfo(false)
	}

  return (
    <div className="page">
      <h1>Eliminar un link corto</h1>
      <LinkInput
        label={'Ingresa la URL corta que quieres eliminar:'}
        placeholder={'http://me.li/XXYYZZ'}
        buttonText={'Eliminar'}
        buttonAction={deleteShortUrl}
        onChangeAction={restartLink}
        loading={loading}
      />
      <br />
          <LinkInfo
            label={label}
            link={{
              url: 'deleted',
              shortUrl: 'deleted',
              createdAt: 'deleted',
            }}
						visible={showLinkInfo}
            showDetails={false}
          />
    </div>
  );
};

export default Delete;
