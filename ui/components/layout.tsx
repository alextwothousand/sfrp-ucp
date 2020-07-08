import Head from 'next/head';
import Navbar from "./navbar";

interface Props {
    children: JSX.Element | JSX.Element[]
}

export default ({ children }) => {
    return (
        <div>
            <Head>
                <title>San Fierro Roleplay</title>
                <meta name="viewport" content="initial-scale=1.0, width=device-width"/>
            </Head>

            <Navbar/>

            <section className="hero is-fullheight-with-navbar">
                <div className="hero-body">
                    <div className="container has-text-centered">
                        <h1 className="title">
                            Fullheight title
                        </h1>
                        <h2 className="subtitle">
                            Fullheight subtitle
                        </h2>
                        {children}
                    </div>
                </div>
            </section>
        </div>
    )
}