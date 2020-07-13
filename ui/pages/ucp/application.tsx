import React from "react";
import Auth from "../../utils/auth";
import Link from "next/link";
import Ucp from "../../components/ucp";
import { NextPageContext } from "next";

class Application extends React.Component {
	async getStaticProps(ctx: NextPageContext): Promise<unknown> {
		await Auth(ctx);
		return {};
	}

	render(): JSX.Element {
		return (
			<Ucp title="UCP - Application" active={10} hascharacters={false} application={false} sidebar={true} navbar={true}>
				<h1>You have successfully logged in! Well done mate. Yes. AAAAA.</h1>
				<h2><Link href="/logout"><a>Logout Here</a></Link></h2>
			</Ucp>
		);
	}
}

export default Application;