//go:build ignore

///*
//Copyright © 2022 NAME HERE <EMAIL ADDRESS>
//
//*/
//package cmd
//
//import (
//	"fmt"
//
//	"github.com/spf13/cobra"
//)
//
//// queryCmd represents the query command
//var queryCmd = &cobra.Command{
//	Use:   "query",
//	Short: "Get Query",
//	Long: `### Get a previously created query by id.
//
//A Looker query object includes the various parameters that define a database query that has been run or
//could be run in the future. These parameters include: model, view, fields, filters, pivots, etc.
//Query *results* are not part of the query object.
//
//Query objects are unique and immutable. Query objects are created automatically in Looker as users explore data.
//Looker does not delete them; they become part of the query history. When asked to create a query for
//any given set of parameters, Looker will first try to find an existing query object with matching
//parameters and will only create a new object when an appropriate object can not be found.
//
//This 'get' method is used to get the details about a query for a given id. See the other methods here
//to 'create' and 'run' queries.
//
//Note that some fields like 'filter_config' and 'vis_config' etc are specific to how the Looker UI
//builds queries and visualizations and are not generally useful for API use. They are not required when
//creating new queries and can usually just be ignored.`,
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("query called")
//	},
//}
//
//func init() {
//	rootCmd.AddCommand(queryCmd)
//
//	// Here you will define your flags and configuration settings.
//
//	// Cobra supports Persistent Flags which will work for this command
//	// and all subcommands, e.g.:
//	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")
//
//	// Cobra supports local flags which will only run when this command
//	// is called directly, e.g.:
//	queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
//}
