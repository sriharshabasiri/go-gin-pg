<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Sample HTML Email</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            color: #333;
        }
        h1 {
            color: #4CAF50;
        }
        table {
            width: 100%;
            border-collapse: collapse;
        }
        th, td {
            padding: 10px;
            text-align: left;
            border: 1px solid #ddd;
        }
        th {
            background-color: #f2f2f2;
        }
        tr:nth-child(even) {
            background-color: #f9f9f9;
        }
        .success {
            color: green;
        }
        .failure {
            color: red;
        }
        .slno {
            width: 10%;
        }
        .service {
            width: 40%;
        }
        .status {
            width: 50%;
        }
    </style>
</head>
<body>
    <h1>Service Status Report</h1>
    <p>This is a sample HTML email with a dynamically populated status table. Below is the table:</p>
    <table>
        <tr>
            <th class="slno">Sl. No.</th>
            <th class="service">Service</th>
            <th class="status">Status</th>
        </tr>
        <tr>
            <td>1</td>
            <td>Test</td>
            <td class="[[SERVICE1_STATUS_CLASS]]">[[SERVICE1_STATUS]]</td> <!-- Placeholder for the status and color -->
        </tr>
        <tr>
            <td>2</td>
            <td>Test</td>
            <td class="[[SERVICE2_STATUS_CLASS]]">[[SERVICE2_STATUS]]</td> <!-- Placeholder for the status and color -->
        </tr>
        <tr>
            <td>3</td>
            <td>Test</td>
            <td class="[[SERVICE3_STATUS_CLASS]]">[[SERVICE3_STATUS]]</td> <!-- Placeholder for the status and color -->
        </tr>
    </table>
    <p>Thank you for checking the service status!</p>
</body>
</html>
import React, { useState } from 'react';

const FileDeployList = () => {
  const [filename, setFilename] = useState('');
  const [patchList, setPatchList] = useState([]);
  const [error, setError] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch(`http://localhost:8080/patch-list?filename=${filename}`);
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const data = await response.json();
      setPatchList(data);
      setError(null);
    } catch (error) {
      setError('Failed to fetch patch list.');
      console.error(error);
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <label>
          Filename:
          <input type="text" value={filename} onChange={(e) => setFilename(e.target.value)} />
        </label>
        <button type="submit">Get File Deploy List</button>
      </form>

      {error && <p style={{ color: 'red' }}>{error}</p>}

      {patchList.length > 0 && (
        <div>
          <h2>File Deploy List:</h2>
          <table border="1">
            <thead>
              <tr>
                <th>Deploy File</th>
                <th>Source File Size</th>
                <th>Patch ID</th>
                <th>Release ID</th>
                <th>Modification Time</th>
              </tr>
            </thead>
            <tbody>
              {patchList.map((patch) => (
                <tr key={patch.patch_id}>
                  <td>{patch.deploy_file}</td>
                  <td>{patch.src_file_size}</td>
                  <td>{patch.patch_id}</td>
                  <td>{patch.release_id}</td>
                  <td>{patch.mod_time}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
};

export default FileDeployList;
