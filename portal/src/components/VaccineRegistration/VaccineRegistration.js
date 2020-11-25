import React from 'react';
import VaccineRegistrationForm from '../VaccineRegistrationForm/VaccineRegistrationForm';
import styles from './VaccineRegistration.module.css';

function VaccineRegistration() {
    return(
        <div className={styles['container']}>
            <div className={styles['registration-form']}>
            <VaccineRegistrationForm />
        </div>
        <div className={styles['registration-form']}>
            <p>List of Registered medicines</p>
        </div>
        </div>
    );
}

export default VaccineRegistration;