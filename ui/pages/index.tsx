import React from "react";
import axios from "axios";
import Layout from "../components/layout";
import { GetStaticProps } from "next";
import { parseCookies } from "nookies";

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
		<Layout title="San Fierro Roleplay" navbar={true} bg={true}>
			<h1 className="title">
				San Fierro Roleplay
			</h1>
			<h2 className="subtitle">
				A new era of San Andreas: Multiplayer Roleplay.
			</h2>
		</Layout>
	);
};

export const getStaticProps: GetStaticProps = async (ctx: NextPageContext) => {
	const posts = await axios.get("http://localhost:3001/api/v1/news");


	await axios.post("http://localhost:3001/api/v1/quiz", {}, {
		headers: {
			"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTQ2NzA5NjYsImlkIjoxLCJuYW1lIjoiZGV2bGV4YW5kZXIifQ.IRCUr7eLoF4ZinYv17-AVN-2x9HvMVdeKU0S2iyHlCE"
		}
	});

	return {
		props: {
			posts: posts.data
		},
	};
};

export default Index;