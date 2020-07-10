import React from "react";
import Link from "next/link";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faHome } from "@fortawesome/free-solid-svg-icons";

const Navbar = (): JSX.Element => {
	const toggleStyles = () => {
		document.querySelector("#burger").classList.toggle("is-active");
		document.querySelector("#navbarmenu").classList.toggle("is-active");
	};

	return (
		<nav className="navbar" role="navigation" aria-label="main navigation">
			<div className="container">

				<div className="navbar-brand">
					<a className="navbar-item" href="https://bulma.io">
						<img src="https://bulma.io/images/bulma-logo.png" width="112" height="28"/>
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
									<FontAwesomeIcon icon={faHome} />
								</span>
								<span>Home</span>
							</a>
						</Link>

						<Link href="sfrp.devlexander.com">
							<a className="navbar-item">
								Forum
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
						<div className="navbar-item">
							<div className="buttons">
								<a className="button is-dark">
									<strong>Register</strong>
								</a>
								<a className="button is-light">
									Log in
								</a>
							</div>
						</div>
					</div>
				</div>

			</div>
		</nav>
	);
};

export default Navbar;