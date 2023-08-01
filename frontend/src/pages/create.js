import React, { useState } from 'react';
import LinkInput from '../components/LinkInput';
import axios from "axios";
import LinkInfo from '../components/LinkInfo';

const Create = () => {
	const [link, setLink] = useState({
		url: '',
		shortUrl: '',
		createdAt: ''
	})

	const [loading, setLoading] = useState(false)

	function createShortUrl(url) {
		setLoading(true)
		axios.post(`http://localhost/api/v1/create`, { url: url })
		.then((res) => {
			const body = res.data
			setLink({
				url: body.data.link.url,
				shortUrl: body.data.link.short_url,
				createdAt: body.data.link.created_at
			})			
		})
		.catch((err) => {
			console.error(err)
		})
		.then(() => {
			setLoading(false)
		});
	}

	function restartLink() {
		setLink({
			url: '',
			shortUrl: '',
			createdAt: ''
		})
	}

  return (
		<div className='page'>
			<h1>Crear un nuevo link corto</h1>
			<LinkInput
				label={'Ingresa la URL que quieres acortar'}
				placeholder={
					'https://www.mercadolibre.cl/amazon-echo-dot-5th-gen-with-clock-con-asistente-virtual-alexa-pantalla-integrada-cloud-blue-110v240v/p/MLC19757118'
				}
				buttonText={'Acortar'}
				buttonAction={createShortUrl}
				onChangeAction={restartLink}
				loading={loading}
			/>
			<br/>
			<LinkInfo label={'asdasdasd'} link={link} />
		</div>
    
  );
};

export default Create;
