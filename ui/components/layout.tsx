import React from "react";
import Head from "next/head";
import Navbar from "./navbar";

interface Props {
	children: JSX.Element | JSX.Element[]
	navbar: boolean
	title: string
}

const Layout = (props: Props): JSX.Element => {
	return (
		<div>
			<Head>
				<title>{props.title.length > 0 ? props.title : "San Fierro Roleplay"}</title>
				<meta name="viewport" content="initial-scale=1.0, width=device-width"/>

				<script src="https://kit.fontawesome.com/f3f846f7f5.js" crossOrigin="anonymous"></script>
				<style>{"body, html { background-color: #ecedf0; }"}</style>
			</Head>

			{props.navbar ? <Navbar/> : null}

			<section className={props.navbar ? "hero is-fullheight-with-navbar" : "hero is-fullheight"}>
				<div className="hero-body">
					<div className="container has-text-centered">
						{props.children}
					</div>
				</div>
			</section>
		</div>
	);
};

export default Layout;