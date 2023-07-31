import React from "react";
import {NavLink} from "react-router-dom";
import styles from "../assets/components/Navbar.module.css";

const Link = ({path, name}) => {
  return (
    <NavLink className={(navData) => (navData.isActive ? styles.active : styles.link)} to={path}>
      {name}
    </NavLink>
  )
}

const Navbar = () => {
	return (
		<div className={styles.nav}>
      <div className={styles.menu}>
        {
          [{path: '/', name: 'Home'}, {path: '/about', name: 'About'}, {path: '/contact', name: 'Contact'}].map((link) => {
            return <Link path={link.path} name={link.name} />
          })
        }
      </div>
    </div>
	);
};

export default Navbar;
