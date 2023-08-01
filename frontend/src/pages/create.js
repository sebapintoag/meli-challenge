import React, { useState } from 'react';
import LinkInput from '../components/LinkInput';
import LinkInfo from '../components/LinkInfo';
import axios from "axios";

const Create = () => {
	const [link, setLink] = useState({
		url: '',
		shortUrl: '',
		createdAt: ''
	})

	const [showLinkInfo, setshowLinkInfo] = useState(false);
	const [showDetails, setshowDetails] = useState(false)
	const [loading, setLoading] = useState(false)
	const [label, setLabel] = useState('');

	function createShortUrl(url) {
		setLoading(true)
		axios.post(`http://localhost/api/v1/create`, { url: url })
		.then((res) => {
			console.log(res)
			const body = res.data
			setLink({
				url: body.data.link.url,
				shortUrl: body.data.link.short_url,
				createdAt: body.data.link.created_at
			})
			setshowDetails(true)
			if(res.status == 201) {
				setLabel('URL corta creada');
			} else {
				setLabel(`${url} ya tiene una URL corta`);
			}
			
		})
		.catch((err) => {
			setshowDetails(false)
			setLabel('No se pudo crear la URL corta');
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
			<h1>Crear un nuevo link corto</h1>
			<LinkInput
				label={'Ingresa la URL que quieres acortar:'}
				placeholder={
					'https://www.mercadolibre.cl/amazon-echo-dot-5th-gen-with-clock-con-asistente-virtual-alexa-pantalla-integrada-cloud-blue-110v240v/p/MLC19757118'
				}
				buttonText={'Acortar'}
				buttonAction={createShortUrl}
				onChangeAction={restartLink}
				loading={loading}
			/>
			<br/>
			<LinkInfo label={label} link={link} visible={showLinkInfo} showDetails={showDetails} />
		</div>
    
  );
};

export default Create;
