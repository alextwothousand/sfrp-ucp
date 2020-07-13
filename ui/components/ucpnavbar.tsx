import React from "react";
import Link from "next/link";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faHome } from "@fortawesome/free-solid-svg-icons";
import handleLogout from "../utils/logout";

const UcpNavbar = (): JSX.Element => {
	const toggleStyles = () => {
		document.querySelector("#burger").classList.toggle("is-active");
		document.querySelector("#navbarmenu").classList.toggle("is-active");
	};

	return (
		<nav className="navbar" role="navigation" aria-label="main navigation">
			<div className="container">

				<div className="navbar-brand">
					<a className="navbar-item" href="https://bulma.io">
						<img src="/SFlogo2.png" width="70" height="auto"/>
					</a>

					<a id="burger" onClick={toggleStyles} role="button" className="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarmenu">
						<span aria-hidden={"true"}></span>
						<span aria-hidden={"true"}></span>
						<span aria-hidden={"true"}></span>
					</a>
				</div>

				<div id="navbarmenu" className="navbar-menu">
					<div className="navbar-start">
						<Link href="/">
							<a className="navbar-item">
								<span className="icon">
									<i aria-hidden={"true"} className="fas fa-home"></i>
								</span>
								<span>Home</span>
							</a>
						</Link>

						<Link href="sfrp.devlexander.com">
							<a className="navbar-item">
								<span className="icon">
									<i className="fab fa-forumbee"></i>
								</span>
								<span>Forum</span>
							</a>
						</Link>

						<div className="navbar-item has-dropdown is-hoverable">
							<a className="navbar-link">
								More
							</a>

							<div className="navbar-dropdown">
								<a className="navbar-item">
									About
								</a>
								<a className="navbar-item">
									Jobs
								</a>
								<a className="navbar-item">
									Contact
								</a>
								<hr className="navbar-divider"/>
								<a className="navbar-item">
									Report an issue
								</a>
							</div>
						</div>
					</div>

					<div className="navbar-end">
						<div className="navbar-item has-dropdown is-hoverable">
							<a className="navbar-link">
								<span className="icon">
									<i className="fas fa-user"></i>
								</span>
							</a>

							<div className="navbar-dropdown">
								<a className="navbar-item">
									About
								</a>
							</div>
						</div>

						<div className="navbar-item">
							<div className="buttons">
								<a className="button is-dark" onClick={handleLogout}>
									<strong>Logout</strong>
								</a>
							</div>
						</div>
					</div>
				</div>

			</div>
		</nav>
	);
};

export default UcpNavbar;