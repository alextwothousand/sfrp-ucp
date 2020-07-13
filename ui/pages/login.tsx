import React from "react";
import Layout from "../components/layout";
import Link from "next/link";
import axios from "axios";
import Router from "next/router";
import { setCookie } from "nookies";

class Register extends React.Component {
	constructor(props: unknown) {
		super(props);

		this.state = {
			username: "",
			password: "",
			errorLabel: "",
			errorLabelHidden: true,
			list: []
		};
	}

	onChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
		this.setState({[e.currentTarget.name]: e.currentTarget.value});
	}

	onSubmit = async (event: React.FormEvent<HTMLFormElement>): Promise<void> => {
		event.preventDefault();

		const {
			username,
			password,
			errorLabel
		} = this.state;

		/*axios.get("http://localhost:3001/api/v1/user", {

		})*/

		await axios.get("http://localhost:3001/api/v1/user", {
			params: {
				username: username,
				password: password 
			},
			headers: {
				"Content-Type": "application/json",
				"Accept": "application/json"
			},
			validateStatus: function (status) {
				return status < 500;
			}
		}).then(async (resp) => {
			if (resp.status == 200) {
				await axios.post("http://localhost:3001/api/v1/session", {
					username: username
				}).then((res) => {
					if (res.status == 200) {
						setCookie(null, "token", res.data.token, {
							maxAge: 7200,
							path: "/"
						});
	
						Router.push("/ucp");
					}
				}).catch(() => {
					alert("An error has occured. Please try again later.");
				});
			}
		}).catch(() => {
			alert("Please ensure you've entered information into the boxes.");
		});


	};

	render(): JSX.Element {
		const {
			username,
			password,
			list,
			errorLabel
		} = this.state;

		return (
			<Layout title={"UCP - Register"} navbar={false} bg={false}>
				<div className="columns is-desktop">
					<div className="column"></div>
					<div className="column is-half-desktop">
						<img src="/SFlogo2.png" alt="San Fierro Roleplay"/>

						<div className="card px-6 py-4" style={{borderRadius: "4px"}}>
							<div className="card-content">
								<div className="field">
									<p className="mb-4">
										Welcome back to San Fierro Roleplay.
										<br/>
										Fill in the form below and click on <strong>next</strong> to proceed.
									</p>
								</div>

								<form action="post" onSubmit={this.onSubmit}>
									<div className="field is-horizontal py-1">
										<div className="field-label is-normal">
											<label className="label" htmlFor="username">Username</label>
										</div>
										<div className="field-body">
											<div className="field">
												<p className="control is-expanded has-icons-left">
													<input className="input" type="text" name="username" value={username} placeholder="Username" onChange={this.onChange}/>
													<span className="icon is-small is-left">
														<i className="fas fa-user-tag"></i>
													</span>
												</p>
											</div>
										</div>
									</div>

									<div className="field is-horizontal py-1">
										<div className="field-label is-normal">
											<label className="label" htmlFor="password">Password</label>
										</div>
										<div className="field-body">
											<div className="field">
												<p className="control is-expanded has-icons-left">
													<input className="input" type="password" name="password" value={password} placeholder="Password" onChange={this.onChange}/>
													<span className="icon is-small is-left">
														<i aria-hidden className="fas fa-user-tag"/>
													</span>
												</p>
											</div>
										</div>
									</div>

									<div className="field">
										<div className="control">
											<button className="button is-light">
												Login
											</button>
										</div>
									</div>
								</form>

							</div>
						</div>

						<div className="my-3">
							{"Don't"} have an account? <Link href="/register"><a>Register!</a></Link>
						</div>
						
					</div>
					<div className="column"></div>
				</div>

			</Layout>
		);
	}
}

export default Register;