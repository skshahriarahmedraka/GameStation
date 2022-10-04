import { json } from '@sveltejs/kit';
import axios from 'axios';
import minio from 'minio';
import * as fs from 'fs';

import {v4 as uuidv4} from 'uuid';



const minioClient = new minio.Client({
	endPoint: '127.0.0.1',
	port: 9000,
	useSSL: false,
	accessKey: 'admin',
	secretKey: 'miniosecret'
});




export async function POST({ request }) {
	const data = await request.json();
	// console.log('🚀 ~ file: +server.ts ~ line 45 ~ POST ~ data', data);

	const file = data['img'];

	fs.writeFileSync(`static/avatar.png`, file, 'base64');

	// Generate UUID
	let myuuid = uuidv4();

	// UPLOAD in MINIO
	const f = `static/avatar.png`;
	const metaData = fs.statSync(f);
	minioClient.fPutObject('myimg', `avatar.png`, f, metaData, function (err, etag) {
		return console.log(err, etag); // err should be null
	});

	
	const mydata = {
		Link: '',
		Name: 'avatar.png',
		BucketName: 'myimg'
	};
	let ResData: {
		Link: string;
		Name: string;
		BucketName: string;
	};

	await fetch('http://localhost:8001/imglink', {
		method: 'POST',
		mode: 'no-cors',

		headers: {
			'Content-Type': 'application/json',
			Accept: 'application/json'
		},
		body: JSON.stringify(mydata)
	})
		.then((res) => {
			return res.json();
		})
		.then((data) => {
			console.log('🚀 ~ file: +server.ts ~ line 99 ~ .then ~ data', data);

			ResData = data;
		});

	return json(ResData);
}

