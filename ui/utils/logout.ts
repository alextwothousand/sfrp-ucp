import { destroyCookie } from "nookies";
import Router from "next/router";

const handleLogout = (): void => {
	destroyCookie(null, "token");
	Router.push("/");
};

export default handleLogout;