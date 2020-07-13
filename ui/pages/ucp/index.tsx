import React from "react";
import { NextPageContext } from "next";
import Auth from "../../utils/auth";
import Link from "next/link";
import Ucp from "../../components/ucp";

class Index extends React.Component {
	static async getInitialProps(ctx: NextPageContext): Promise<unknown> {
		await Auth(ctx)
			.then(() => console.log("yep"))
			.catch(() => console.log("nope"));
		return {};
	}

	render(): JSX.Element {
		return (
			<Ucp title="UCP - Home" active={1} hascharacters={false} application={false} sidebar={true} navbar={true}>
				<h1>You have successfully logged in! Well done mate.</h1>
				<h2><Link href="/logout"><a>Logout Here</a></Link></h2>
			</Ucp>
		);
	}
}

export default Index;