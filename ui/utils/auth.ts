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

	await axios.get("http://localhost:3001/api/v1/session", 
		{
			headers: {
				"Authorization": "Bearer " + token
			},
		})
		.then((res) => {
			if (res.status == 200) console.log("authorised");
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
		});
};

export default Auth;