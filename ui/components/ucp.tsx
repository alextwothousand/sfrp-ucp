import React from "react";
import Head from "next/head";
import Sidebar from "./sidebar";
import UcpNavbar from "./ucpnavbar";

interface Props {
	sidebar: boolean
	navbar: boolean
	children: JSX.Element | JSX.Element[]
	title: string
	active: number
	application: boolean
	hascharacters: boolean
}

const Ucp = (props: Props): JSX.Element => (
	<div>
		<Head>
			<title>{props.title.length > 0 ? props.title : "San Fierro Roleplay"}</title>
			<meta name="viewport" content="initial-scale=1.0, width=device-width"/>

			<script src="https://kit.fontawesome.com/f3f846f7f5.js" crossOrigin="anonymous" />
			<style>{"body, html { background-color: #ecedf0; }"}</style>
		</Head>

		{props.navbar ? <UcpNavbar/> : null}

		<section className="section">
			<div className="container">
				<div className="columns">
					<div className="column is-one-third">
						{props.sidebar ? <Sidebar Active={props.active} Application={props.application} HasCharacters={props.hascharacters}/> : null}
					</div>
					<div className="column">
						{props.children}
					</div>
				</div>
			</div>
		</section>

	</div>
);

export default Ucp;