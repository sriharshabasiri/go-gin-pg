{/* Button to toggle the visibility of the Create form */}
      <button onClick={() => setShowCreateForm(!showCreateForm)}>
        {showCreateForm ? 'Cancel' : 'Create'}
      </button>

      {/* Create form */}
      {showCreateForm && (
        <form onSubmit={handleSubmit}>
          <table>
            <tbody>
              <tr>
                <td>Environment:</td>
                <td>
                  <select
                    name="environment"
                    value={newRefresh.environment}
                    onChange={handleInputChange}
                    required
                  >
                    {/* Options for Environment dropdown */}
                    <option value="production">Production</option>
                    <option value="testing">Testing</option>
                    <option value="development">Development</option>
                    {/* Add more options as needed */}
                  </select>
                </td>
              </tr>
              <tr>
                <td>Code Base:</td>
                <td>
                  <input
                    type="text"
                    name="code_base"
                    value={newRefresh.code_base}
                    onChange={handleInputChange}
                    required
                  />
                </td>
              </tr>
              <tr>
                <td>Change Ticket:</td>
                <td>
                  <input
                    type="text"
                    name="change_ticket"
                    value={newRefresh.change_ticket}
                    onChange={handleInputChange}
                    required
                  />
                </td>
              </tr>
            </tbody>
          </table>

          <button type="submit">Submit</button>
        </form>
      )}
