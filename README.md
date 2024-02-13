import React, { useState, useEffect } from 'react';

const RefreshDetails = () => {
  const [refreshes, setRefreshes] = useState([]);
  const [newRefresh, setNewRefresh] = useState({
    environment: '',
    sl_no: '',
    date_of_refresh: '',
    code_base: '',
    change_ticket: '',
    free_field_1: '',
    free_field_2: '',
    free_field_3: '',
    del_flg: false,
    r_mod_time: '',
    r_cre_time: '',
  });

  const fetchRefreshes = async () => {
    try {
      const response = await fetch('your_api_url_for_get_refreshes');
      const data = await response.json();
      setRefreshes(data);
    } catch (error) {
      console.error('Error fetching refreshes:', error);
    }
  };

  const handleInputChange = (e) => {
    const { name, value, type, checked } = e.target;
    // Handle checkbox separately
    const newValue = type === 'checkbox' ? checked : value;

    setNewRefresh((prevRefresh) => ({
      ...prevRefresh,
      [name]: newValue,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('your_api_url_for_post_refresh', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(newRefresh),
      });

      if (response.ok) {
        // Refresh the list after successful submission
        fetchRefreshes();
        setNewRefresh({
          environment: '',
          sl_no: '',
          date_of_refresh: '',
          code_base: '',
          change_ticket: '',
          free_field_1: '',
          free_field_2: '',
          free_field_3: '',
          del_flg: false,
          r_mod_time: '',
          r_cre_time: '',
        });
      } else {
        console.error('Failed to submit refresh details');
      }
    } catch (error) {
      console.error('Error submitting refresh details:', error);
    }
  };

  useEffect(() => {
    // Fetch refreshes when the component mounts
    fetchRefreshes();
  }, []);

  return (
    <div>
      <h2>Refresh Details</h2>

      {/* Display the list of refreshes in a table */}
      <table>
        <thead>
          <tr>
            <th>Environment</th>
            <th>SL No</th>
            <th>Date of Refresh</th>
            <th>Code Base</th>
            <th>Change Ticket</th>
            <th>Free Field 1</th>
            <th>Free Field 2</th>
            <th>Free Field 3</th>
            <th>Delete Flag</th>
            <th>R Modification Time</th>
            <th>R Creation Time</th>
          </tr>
        </thead>
        <tbody>
          {refreshes.map((refresh) => (
            <tr key={refresh.sl_no}>
              {/* Display relevant details in each column */}
              <td>{refresh.environment}</td>
              <td>{refresh.sl_no}</td>
              <td>{refresh.date_of_refresh}</td>
              <td>{refresh.code_base}</td>
              <td>{refresh.change_ticket}</td>
              <td>{refresh.free_field_1}</td>
              <td>{refresh.free_field_2}</td>
              <td>{refresh.free_field_3}</td>
              <td>{refresh.del_flg ? 'Yes' : 'No'}</td>
              <td>{refresh.r_mod_time}</td>
              <td>{refresh.r_cre_time}</td>
            </tr>
          ))}
        </tbody>
      </table>

      {/* Form for submitting new refresh details */}
      <form onSubmit={handleSubmit}>
        {/* ... (unchanged) */}
      </form>
    </div>
  );
};

export default RefreshDetails;
