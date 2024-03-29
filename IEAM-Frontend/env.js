const envCf = require('dotenv');
envCf.config();

let FIREBASE = {	
	apiKey: "AIzaSyANHmZOXv6xf0T3pLX8IDvrWbi-HVy3xq4",
    authDomain: "blog-5e5ba.firebaseapp.com",
    projectId: "blog-5e5ba",
    storageBucket: "blog-5e5ba.appspot.com",
    messagingSenderId: "135513270144",
    appId: "1:135513270144:web:60be2256d54496af1ac365",
    measurementId: "G-R11W3XST73"
}

let env = {
	development: {
		VERSION: '1.0.0(beta)',
		BASE_URL:'https://ieam-api.ichigozdata.win',
		FIREBASE,
	},
	production: {
		VERSION: '1.0.1',
		BASE_URL:'https://ieam-api.ichigozdata.win',
		FIREBASE,
	},
}
console.log(env[process.env.API_ENV]);
module.exports = env[process.env.API_ENV]
    ? env[process.env.API_ENV]
    : env["development"];
