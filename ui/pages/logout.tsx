import React from "react";
import Layout from "../components/layout";
import { NextPageContext } from "next";
import Auth from "../utils/auth";
import Link from "next/link";
import axios from "axios";

class Logout extends React.Component {
	async getInitialProps(ctx: NextPageContext): Promise<unknown> {
		await Auth(ctx);
		return {};
	}

	async render (): JSX.Element {
		await axios.delete("");

		return (
			<Layout title="UCP - Home" navbar={false}>
				<h1>You have successfully logged in! Well done mate.</h1>
				<h2><Link href="/logout"><a>Logout Here</a></Link></h2>
			</Layout>
		);
	}
}


export default Logout;