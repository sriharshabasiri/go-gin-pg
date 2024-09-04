#!/bin/bash

# Define email details
SUBJECT="Service Status Report"
FROM_EMAIL="sender@example.com"
TO_EMAIL="recipient@example.com"

# Define status values for the services
SERVICE1_STATUS="$1"  # Pass as argument or set directly (e.g., "Success" or "Failure")
SERVICE2_STATUS="$2"  # Pass as argument or set directly
SERVICE3_STATUS="$3"  # Pass as argument or set directly

# Determine CSS class based on status
SERVICE1_STATUS_CLASS=$( [[ "$SERVICE1_STATUS" == "Success" ]] && echo "success" || echo "failure" )
SERVICE2_STATUS_CLASS=$( [[ "$SERVICE2_STATUS" == "Success" ]] && echo "success" || echo "failure" )
SERVICE3_STATUS_CLASS=$( [[ "$SERVICE3_STATUS" == "Success" ]] && echo "success" || echo "failure" )

# Load the HTML template and replace placeholders
TEMPLATE_PATH="email_template.html"
OUTPUT_PATH="email.html"

# Replace placeholders in the template
sed -e "s/\[\[SERVICE1_STATUS\]\]/$SERVICE1_STATUS/g" \
    -e "s/\[\[SERVICE1_STATUS_CLASS\]\]/$SERVICE1_STATUS_CLASS/g" \
    -e "s/\[\[SERVICE2_STATUS\]\]/$SERVICE2_STATUS/g" \
    -e "s/\[\[SERVICE2_STATUS_CLASS\]\]/$SERVICE2_STATUS_CLASS/g" \
    -e "s/\[\[SERVICE3_STATUS\]\]/$SERVICE3_STATUS/g" \
    -e "s/\[\[SERVICE3_STATUS_CLASS\]\]/$SERVICE3_STATUS_CLASS/g" \
    "$TEMPLATE_PATH" > "$OUTPUT_PATH"

# Define the path to the Python script
PYTHON_SCRIPT_PATH="/path/to/send_email.py"  # Update this to the correct path

# Check if the Python script exists
if [[ -f "$PYTHON_SCRIPT_PATH" ]]; then
    echo "Running the Python script to send the email..."
    
    # Run the Python script with arguments
    python3 "$PYTHON_SCRIPT_PATH" --subject "$SUBJECT" --from_email "$FROM_EMAIL" --to_email "$TO_EMAIL"
else
    echo "Python script not found at $PYTHON_SCRIPT_PATH"
    exit 1
fi
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
