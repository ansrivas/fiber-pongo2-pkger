//
// Copyright (c) 2020 Ankur Srivastava
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package routers

import "github.com/gofiber/fiber/v2"

// Response being sent out from the server
type Response struct {
	Status  int
	Message string
	Data    interface{}
}

// SendJSON is a wrapper over exisiting json response of fiber
func SendJSON(c *fiber.Ctx, statusCode int, message string, data interface{}) error {

	var resp interface{}
	switch val := data.(type) {

	// Some concrete type in case you want to test against
	// case something-special:

	default:
		resp = Response{
			Status:  statusCode,
			Message: message,
			Data:    val,
		}
	}

	if err := c.Status(statusCode).JSON(resp); err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}
	return nil
}
