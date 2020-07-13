import axios from "axios";
import Router from "next/router";
import { parseCookies } from "nookies";
import { NextPageContext } from "next";

const Auth = async (ctx: NextPageContext): Promise<void> => {
	let token = null;

	if (ctx.req) {
		token = ctx.req.headers.cookie.replace(/(?:(?:^|.*;\s*)token\s*=\s*([^;]*).*$)|^.*$/, "$1");
	} else {
		token = parseCookies(ctx).token;
	}

	console.log("token is " + token);

	await axios.get("http://localhost:3001/api/v1/session", 
		{
			headers: {
				"Authorization": "Bearer " + token
			},
		})
		.then((res) => {
			console.log(res.status);
			if (res.status == 200) Promise.resolve();
		})
		.catch(() => {
			if (ctx.res) {
				ctx.res.writeHead(302, {
					Location: "/"
				});

				ctx.res.end();
			} else {
				Router.push("/");
			}

			Promise.reject();
		});
};

export default Auth;