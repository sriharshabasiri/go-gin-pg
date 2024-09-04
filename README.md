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
    </style>
</head>
<body>
    <h1>Welcome to the HTML Email Test!</h1>
    <p>This is a sample HTML email with some basic formatting, a table, and styled text. Below is a simple table:</p>
    <table>
        <tr>
            <th>Item</th>
            <th>Quantity</th>
            <th>Price</th>
        </tr>
        <tr>
            <td>Apples</td>
            <td>10</td>
            <td>$5.00</td>
        </tr>
        <tr>
            <td>Bananas</td>
            <td>5</td>
            <td>$2.50</td>
        </tr>
        <tr>
            <td>Oranges</td>
            <td>8</td>
            <td>$4.00</td>
        </tr>
    </table>
    <p>Thank you for testing HTML emails!</p>
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
