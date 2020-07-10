import React from "react";
import axios from "axios";
import Layout from "../components/layout";
import { GetStaticProps } from "next";

/*interface Props {
	posts: [
		{
			Author: string,
			Text: string
		}
	]
}


			<section className="section">
				<div className="container">
					<div className="card">
						<div className="card-content">
							{props.posts.map((post, i) => (
								<div key={i}>
									<p>{post.Author}</p>
									<p>{post.Text}</p>
								</div>
							))}
						</div>
					</div>
				</div>
			</section>
*/

const Index = (): JSX.Element => {
	return (
		<Layout title="San Fierro Roleplay" navbar={true}>
			<h1 className="title">
				San Fierro Roleplay
			</h1>
			<h2 className="subtitle">
				A new era of San Andreas: Multiplayer Roleplay.
			</h2>
		</Layout>
	);
};

export const getStaticProps: GetStaticProps = async () => {
	const posts = await axios.get("http://localhost:3001/api/v1/news");

	return {
		props: {
			posts: posts.data
		},
	};
};

export default Index;