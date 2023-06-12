import React, { useState, useEffect } from 'react';
import axios from 'axios';

const PatientList = () => {
  const [patients, setPatients] = useState([]);
  const [newPatient, setNewPatient] = useState({
    id: '',
    first_name: '',
    last_name: '',
    date_of_birth: '',
    gender: '',
    email: '',
    phone_number: '',
    address: '',
    created_at: '2023-06-06',
  });

  useEffect(() => {
    fetchPatients();
  }, []);

  const fetchPatients = async () => {
    try {
      const response = await axios.get('http://localhost:8080/patients');
      setPatients(response.data);
    } catch (error) {
      console.log(error);
    }
  };

  const handleInputChange = (e) => {
    setNewPatient({
      ...newPatient,
      [e.target.name]: e.target.value,
    });
  };

  const addPatient = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post(
        'http://localhost:8080/patients',
        {
          id: parseInt(newPatient.id),
          first_name: newPatient.first_name,
          last_name: newPatient.last_name,
          date_of_birth: newPatient.date_of_birth,
          gender: newPatient.gender,
          email: newPatient.email,
          phone_number: newPatient.phone_number,
          address: newPatient.address,
          created_at: new Date().toISOString().split('T')[0],
        },
        {
          headers: {
            'Content-Type': 'application/json',
          },
        }
      );

      console.log(response.data);

      // Add the current response to the list of already existing patient data
      setPatients([...patients, response.data]);

      // Reset the newPatient state
      setNewPatient({
        id: '',
        first_name: '',
        last_name: '',
        date_of_birth: '',
        gender: '',
        email: '',
        phone_number: '',
        address: '',
        created_at: new Date().toISOString().split('T')[0],
      });
    } catch (error) {
      console.log(error);
    }
  };

  const deletePatient = async (id) => {
    try {
      await axios.delete(`http://localhost:8080/patients/${id}`);
      const updatedPatients = patients.filter((patient) => patient.id !== id);
      setPatients(updatedPatients);
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div style={{ display: 'flex' }}>
      <div style={{ width: '70%' }}>
        <h2>Patients List</h2>
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>First Name</th>
              <th>Last Name</th>
              <th>Date of Birth</th>
              <th>Gender</th>
              <th>Email</th>
              <th>Phone Number</th>
              <th>Address</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {patients.map((patient) => (
              <tr key={patient.id}>
                <td>{patient.id}</td>
                <td>{patient.first_name}</td>
                <td>{patient.last_name}</td>
                <td>{patient.date_of_birth}</td>
                <td>{patient.gender}</td>
                <td>{patient.email}</td>
                <td>{patient.phone_number}</td>
                <td>{patient.address}</td>
                <td>
                  <button onClick={() => deletePatient(patient.id)}>Delete</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <div style={{ width: '30%', padding: '20px' }}>
        <h2>Add New Patient</h2>
        <form onSubmit={addPatient} style={{ display: 'grid', gap: '10px' }}>
          <div>
            <label style={{ display: 'inline-block', width: '100px' }}>ID:</label>
            <input
              type="text"
              name="id"
              value={newPatient.id}
              onChange={handleInputChange}
            />
          </div>
          <div>
            <label style={{ display: 'inline-block', width: '100px' }}>First Name:</label>
            <input
              type="text"
              name="first_name"
              value={newPatient.first_name}
              onChange={handleInputChange}
            />
          </div>
          <div>
            <label style={{ display: 'inline-block', width: '100px' }}>Last Name:</label>
            <input
              type="text"
              name="last_name"
              value={newPatient.last_name}
              onChange={handleInputChange}
            />
          </div>
          <div>
            <label style={{ display: 'inline-block', width: '100px' }}>Date of Birth:</label>
            <input
              type="text"
              name="date_of_birth"
              value={newPatient.date_of_birth}
              onChange={handleInputChange}
            />
          </div>
          <div>
            <label style={{ display: 'inline-block', width: '100px' }}>Gender:</label>
            <input
              type="text"
              name="gender"
              value={newPatient.gender}
              onChange={handleInputChange}
            />
          </div>
          <div>
            <label style={{ display: 'inline-block', width: '100px' }}>Email:</label>
            <input
              type="text"
              name="email"
              value={newPatient.email}
              onChange={handleInputChange}
            />
          </div>
          <div>
            <label style={{ display: 'inline-block', width: '100px' }}>Phone Number:</label>
            <input
              type="text"
              name="phone_number"
              value={newPatient.phone_number}
              onChange={handleInputChange}
            />
          </div>
          <div>
            <label style={{ display: 'inline-block', width: '100px' }}>Address:</label>
            <input
              type="text"
              name="address"
              value={newPatient.address}
              onChange={handleInputChange}
            />
          </div>
          <div>
            <button type="submit">Add Patient</button>
          </div>
        </form>
      </div>

    </div>
  );
};

export default PatientList;
