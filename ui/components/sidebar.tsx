import React from "react";
import Link from "next/link";

/*
	Available __Active__ Tabs:

	[General]
	1 - Profile
	2 - Your Characters
		3 - Char 1
		4 - Char 2
		5 - Char 3
		6 - Char 4
	7 - Create a Character (Already has chars)
	8 - Delete a Character (Already has chars)
	9 - Create a Character (doesnt have chars)
	10 - Submit Application (doesnt have chars)

	[Administration]
	11 - View Applications
	12 - View Player
	13 - Modify Player
	14 - Modify Character
*/

interface Props {
	HasCharacters: boolean // True if player has characters. False otherwise. Ignored if they don't have an application. 
	Application: boolean // True if application is completed. False otherwise.
	Active: number // Tab that is currently active.
}

const Sidebar = (props: Props): JSX.Element => (
	<aside className="menu">
		<p className="menu-label">
			General
		</p>
		<ul className="menu-list">
			<li><Link href="/ucp"><a className={props.Active == 1 ? "is-active" : ""}>Your Profile</a></Link></li>
			{props.Application ? (
				props.HasCharacters ? (
					<div>
						<li><a className={props.Active == 2 ? "is-active" : ""}>Your Characters</a></li>
						<li>
							<ul>
								<li><a className={props.Active == 3 ? "is-active" : ""}>John_Doe</a></li>
								<li><a className={props.Active == 4 ? "is-active" : ""}>Jane_Doe</a></li>
								<li><a className={props.Active == 5 ? "is-active" : ""}>Sam_Doe</a></li>
								<li><a className={props.Active == 6 ? "is-active" : ""}>Jackie_Doe</a></li>
							</ul>
						</li>
						<li><a className={props.Active == 7 ? "is-active" : ""}>Create a Character</a></li>
						<li><a className={props.Active == 8 ? "is-active" : ""}>Delete a Character</a></li>
					</div>
				) : (
					<li><a className={props.Active == 9 ? "is-active" : ""}>Create a Character</a></li>
				)
			) : (
				<li><Link href="/ucp/application"><a className={props.Active == 10 ? "is-active" : ""}>Submit Application</a></Link></li>
				
			)}
		</ul>
		<p className="menu-label">
			Administration
		</p>
		<ul className="menu-list">
			<li><a className={props.Active == 11 ? "is-active" : ""}>View Applications</a></li>
			<li><a className={props.Active == 12 ? "is-active" : ""}>View Player</a></li>
			<li><a className={props.Active == 13 ? "is-active" : ""}>Modify Player</a></li>
			<li><a className={props.Active == 14 ? "is-active" : ""}>Modify Character</a></li>
		</ul>
	</aside>
);

export default Sidebar;