using System;
using System.Collections.Generic;
using System.Configuration;
using System.Data;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using MySql.Data.MySqlClient;

namespace CSharp.Controllers
{
    public class TextController : Controller
    {
        private const string mySQL = "server=192.168.29.113;database=MONEYSAVERDB;user=suser;password=1532";

        [HttpGet]
        [Route("text")]
        public IActionResult GetText()
        {
            string message = "hello world";
            return Json(message);
        }

        [HttpGet]
        [Route("email")]
        public IActionResult GetEmail(int id)
        {
            string email = "";
            using(MySqlConnection connection = new MySqlConnection("server=192.168.29.113;database=databasename;user=username;password=password"))
            {
                connection.Open();
                MySqlCommand cmd = new MySqlCommand("select Email from User where UserId = "+ id, connection);
                using (var reader = cmd.ExecuteReader())
                {
                    if(reader.HasRows)
                    
                    {
                        reader.Read();
                        email = Convert.ToString(reader["Email"]);
                    }
                }
            }
                return Json(email);
        }
    }
}
