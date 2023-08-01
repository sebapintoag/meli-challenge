import React, { useState } from 'react';
import LinkInput from '../components/LinkInput';
import LinkInfo from '../components/LinkInfo';
import axios from "axios";

const Find = () => {
	const [link, setLink] = useState({
		url: '',
		shortUrl: '',
		createdAt: ''
	})

	const [showLinkInfo, setshowLinkInfo] = useState(false);
	const [showDetails, setshowDetails] = useState(false)
	const [loading, setLoading] = useState(false)
	const [label, setLabel] = useState('');

	function findShortUrl(short_url) {
		setLoading(true)
		axios.post(`http://localhost/api/v1/find`, { short_url: short_url })
		.then((res) => {
			const body = res.data
			setLink({
				url: body.data.link.url,
				shortUrl: body.data.link.short_url,
				createdAt: body.data.link.created_at
			})
			setshowDetails(true)
			setLabel('URL corta encontrada');
		})
		.catch((err) => {
			setshowDetails(false)
			setLabel('No se encontró la URL corta');
		})
		.then(() => {
			setshowLinkInfo(true)
			setLoading(false)
		});
	}

	function restartLink() {
		setshowLinkInfo(false)
	}

  return (
		<div className='page'>
			<h1>Buscar un link corto</h1>
			<LinkInput
				label={'Ingresa la URL corta que quieres buscar:'}
				placeholder={
					'http://me.li/XXYYZZ'
				}
				buttonText={'Buscar'}
				buttonAction={findShortUrl}
				onChangeAction={restartLink}
				loading={loading}
			/>
			<br/>
			<LinkInfo label={label} link={link} visible={showLinkInfo} showDetails={showDetails} />
		</div>
    
  );
};

export default Find;
