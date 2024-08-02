<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>HTML Table Example</title>
    <style>
        table {
            width: 100%;
            border-collapse: collapse;
        }
        th, td {
            border: 1px solid black;
            padding: 8px;
            text-align: left;
        }
        th {
            background-color: #f2f2f2;
        }
        .sl-no {
            width: 2ch;
        }
        .static-text {
            width: 60ch;
        }
        .dynamic-text {
            width: 30ch;
        }
        .green {
            color: green;
        }
        .red {
            color: red;
        }
    </style>
</head>
<body>
    <table>
        <tr>
            <th class="sl-no">Sl No</th>
            <th class="static-text">Static Text</th>
            <th class="dynamic-text">Dynamic Text</th>
        </tr>
        <tr>
            <td>1</td>
            <td>Static Text 1</td>
            <td class="green">Dynamic Green 1</td>
        </tr>
        <tr>
            <td>2</td>
            <td>Static Text 2</td>
            <td class="red">Dynamic Red 2</td>
        </tr>
        <tr>
            <td>3</td>
            <td>Static Text 3</td>
            <td class="green">Dynamic Green 3</td>
        </tr>
        <tr>
            <td>4</td>
            <td>Static Text 4</td>
            <td class="red">Dynamic Red 4</td>
        </tr>
        <tr>
            <td>5</td>
            <td>Static Text 5</td>
            <td class="green">Dynamic Green 5</td>
        </tr>
        <tr>
            <td>6</td>
            <td>Static Text 6</td>
            <td class="red">Dynamic Red 6</td>
        </tr>
        <tr>
            <td>7</td>
            <td>Static Text 7</td>
            <td class="green">Dynamic Green 7</td>
        </tr>
        <tr>
            <td>8</td>
            <td>Static Text 8</td>
            <td class="red">Dynamic Red 8</td>
        </tr>
        <tr>
            <td>9</td>
            <td>Static Text 9</td>
            <td class="green">Dynamic Green 9</td>
        </tr>
        <tr>
            <td>10</td>
            <td>Static Text 10</td>
            <td class="red">Dynamic Red 10</td>
        </tr>
    </table>
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
