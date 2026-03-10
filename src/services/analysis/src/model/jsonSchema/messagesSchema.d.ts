export interface MySchema {
  messages?: (
    | {
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        created_date?: string;
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          url_metadata?: {
            image_url?: string;
            snippet?: string;
            title?: string;
            url?: {
              private_do_not_access_or_else_safe_url_wrapped_value?: string;
            };
          };
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          url_metadata?: {
            image_url?: string;
            snippet?: string;
            title?: string;
            url?: {
              private_do_not_access_or_else_safe_url_wrapped_value?: string;
            };
          };
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          drive_metadata?: {
            id?: string;
            thumbnail_url?: string;
            title?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          url_metadata?: {
            image_url?: string;
            snippet?: string;
            title?: string;
            url?: {
              private_do_not_access_or_else_safe_url_wrapped_value?: string;
            };
          };
        }[];
        created_date?: string;
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          youtube_metadata?: {
            id?: string;
            start_time?: number;
          };
        }[];
        created_date?: string;
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          youtube_metadata?: {
            id?: string;
            start_time?: number;
          };
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          url_metadata?: {
            image_url?: string;
            snippet?: string;
            title?: string;
            url?: {
              private_do_not_access_or_else_safe_url_wrapped_value?: string;
            };
          };
        }[];
        created_date?: string;
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          drive_metadata?: {
            id?: string;
            thumbnail_url?: string;
            title?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: (
          | {
              format_metadata?: {
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
          | {
              length?: number;
              start_index?: number;
              url_metadata?: {
                image_url?: string;
                snippet?: string;
                title?: string;
                url?: {
                  private_do_not_access_or_else_safe_url_wrapped_value?: string;
                };
              };
            }
        )[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: (
          | {
              length?: number;
              start_index?: number;
              url_metadata?: {
                image_url?: string;
                snippet?: string;
                title?: string;
                url?: {
                  private_do_not_access_or_else_safe_url_wrapped_value?: string;
                };
              };
            }
          | {
              format_metadata?: {
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
        )[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: (
          | {
              drive_metadata?: {
                id?: string;
                thumbnail_url?: string;
                title?: string;
              };
              length?: number;
              start_index?: number;
            }
          | {
              format_metadata?: {
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
        )[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: (
          | {
              format_metadata?: {
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
          | {
              length?: number;
              start_index?: number;
              youtube_metadata?: {
                id?: string;
                start_time?: number;
              };
            }
        )[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: (
          | {
              format_metadata?: {
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
          | {
              length?: number;
              start_index?: number;
              url_metadata?: {
                image_url?: string;
                snippet?: string;
                title?: string;
                url?: {
                  private_do_not_access_or_else_safe_url_wrapped_value?: string;
                };
              };
            }
        )[];
        created_date?: string;
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: (
          | {
              format_metadata?: {
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
          | {
              length?: number;
              start_index?: number;
              youtube_metadata?: {
                id?: string;
                start_time?: number;
              };
            }
        )[];
        created_date?: string;
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          gsuite_integration_metadata?: {
            tasks_data?: {
              creation?: {};
              task_properties?: {
                completed?: boolean;
                deleted?: boolean;
                description?: string;
                title?: string;
              };
            };
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          video_call_metadata?: {
            meeting_space?: {
              meeting_url?: string;
            };
          };
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          drive_metadata?: {
            id?: string;
            thumbnail_url?: string;
            title?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          gsuite_integration_metadata?: {
            tasks_data?: {
              completion_change?: {};
              task_properties?: {
                completed?: boolean;
                deleted?: boolean;
                description?: string;
                title?: string;
              };
            };
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          attached_files?: {
            export_name?: string;
            original_name?: string;
          }[];
          created_date?: string;
        }[];
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          video_call_metadata?: {
            meeting_space?: {
              meeting_url?: string;
            };
          };
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        topic_id?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          gsuite_integration_metadata?: {
            calendar_event_data?: {
              calendar_event?: {
                end_time?: {
                  timed?: string;
                };
                start_time?: {
                  timed?: string;
                };
                title?: string;
              };
            };
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: (
          | {
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              updated_date?: string;
            }
          | {
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              created_date?: string;
            }
        )[];
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          url_metadata?: {
            image_url?: string;
            snippet?: string;
            title?: string;
            url?: {
              private_do_not_access_or_else_safe_url_wrapped_value?: string;
            };
          };
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          attached_files?: {
            export_name?: string;
            original_name?: string;
          }[];
          created_date?: string;
          text?: string;
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          drive_metadata?: {
            id?: string;
            thumbnail_url?: string;
            title?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          attached_files?: {
            export_name?: string;
            original_name?: string;
          }[];
          created_date?: string;
        }[];
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            font_color?: number;
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          attached_files?: {
            export_name?: string;
            original_name?: string;
          }[];
          created_date?: string;
          text?: string;
        }[];
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: (
          | {
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              text?: string;
              updated_date?: string;
            }
          | {
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              created_date?: string;
              text?: string;
            }
        )[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        created_date?: string;
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          gsuite_integration_metadata?: {
            tasks_data?: {
              assignee_change?: {
                old_assignee?: {
                  id?: string;
                };
              };
              task_properties?: {
                assignee?: {
                  id?: string;
                };
                completed?: boolean;
                deleted?: boolean;
                description?: string;
                title?: string;
              };
            };
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          youtube_metadata?: {
            id?: string;
            start_time?: number;
          };
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          attached_files?: {
            export_name?: string;
            original_name?: string;
          }[];
          created_date?: string;
        }[];
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: (
          | {
              format_metadata?: {
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
          | {
              format_metadata?: {
                font_color?: number;
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
        )[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: (
          | {
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              updated_date?: string;
            }
          | {
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              created_date?: string;
            }
        )[];
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        topic_id?: string;
        updated_date?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: (
          | {
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              updated_date?: string;
            }
          | {
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              created_date?: string;
            }
        )[];
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          gsuite_integration_metadata?: {
            calendar_event_data?: {
              calendar_event?: {
                end_time?: {
                  timed?: string;
                };
                start_time?: {
                  timed?: string;
                };
                title?: string;
              };
            };
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          acting_user?: {
            name?: string;
            user_id?: string;
            user_type?: string;
          };
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        topic_id?: string;
      }
    | {
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          created_date?: string;
          text?: string;
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: (
          | {
              text?: string;
              updated_date?: string;
            }
          | {
              created_date?: string;
              text?: string;
            }
        )[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
      }
    | {
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          created_date?: string;
          text?: string;
        }[];
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        topic_id?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          annotations?: {
            format_metadata?: {
              format_type?: string;
            };
            length?: number;
            start_index?: number;
          }[];
          created_date?: string;
          text?: string;
        }[];
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: (
          | {
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              updated_date?: string;
            }
          | {
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              created_date?: string;
            }
        )[];
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          url_metadata?: {
            image_url?: string;
            snippet?: string;
            title?: string;
            url?: {
              private_do_not_access_or_else_safe_url_wrapped_value?: string;
            };
          };
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          annotations?: {
            length?: number;
            start_index?: number;
            url_metadata?: {
              image_url?: string;
              snippet?: string;
              title?: string;
              url?: {
                private_do_not_access_or_else_safe_url_wrapped_value?: string;
              };
            };
          }[];
          created_date?: string;
          text?: string;
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          annotations?: {
            length?: number;
            start_index?: number;
            url_metadata?: {
              image_url?: string;
              snippet?: string;
              title?: string;
              url?: {
                private_do_not_access_or_else_safe_url_wrapped_value?: string;
              };
            };
          }[];
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
      }
    | {
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          attached_files?: {
            export_name?: string;
            original_name?: string;
          }[];
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: (
          | {
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              quoted_message_metadata?: {
                creator?: {
                  name?: string;
                  user_type?: string;
                };
                text?: string;
              };
              text?: string;
              updated_date?: string;
            }
          | {
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              created_date?: string;
              quoted_message_metadata?: {
                creator?: {
                  name?: string;
                  user_type?: string;
                };
                text?: string;
              };
              text?: string;
            }
        )[];
        quoted_message_metadata?: {
          creator?: {
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          created_date?: string;
          text?: string;
        }[];
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          attached_files?: {
            export_name?: string;
            original_name?: string;
          }[];
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          url_metadata?: {
            image_url?: string;
            snippet?: string;
            title?: string;
            url?: {
              private_do_not_access_or_else_safe_url_wrapped_value?: string;
            };
          };
        }[];
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          created_date?: string;
          quoted_message_metadata?: {
            creator?: {
              email?: string;
              name?: string;
              user_type?: string;
            };
            text?: string;
          };
          text?: string;
        }[];
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          attached_files?: {
            export_name?: string;
            original_name?: string;
          }[];
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: (
          | {
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              text?: string;
              updated_date?: string;
            }
          | {
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              created_date?: string;
              text?: string;
            }
        )[];
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          annotations?: {
            length?: number;
            start_index?: number;
            url_metadata?: {
              image_url?: string;
              snippet?: string;
              title?: string;
              url?: {
                private_do_not_access_or_else_safe_url_wrapped_value?: string;
              };
            };
          }[];
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          annotations?: {
            length?: number;
            start_index?: number;
            url_metadata?: {
              image_url?: string;
              snippet?: string;
              title?: string;
              url?: {
                private_do_not_access_or_else_safe_url_wrapped_value?: string;
              };
            };
          }[];
          created_date?: string;
          text?: string;
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          annotations?: {
            format_metadata?: {
              format_type?: string;
            };
            length?: number;
            start_index?: number;
          }[];
          created_date?: string;
          text?: string;
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          youtube_metadata?: {
            id?: string;
            start_time?: number;
          };
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          annotations?: {
            length?: number;
            start_index?: number;
            youtube_metadata?: {
              id?: string;
              start_time?: number;
            };
          }[];
          created_date?: string;
          text?: string;
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          annotations?: {
            format_metadata?: {
              format_type?: string;
            };
            length?: number;
            start_index?: number;
          }[];
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
      }
    | {
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: (
          | {
              text?: string;
              updated_date?: string;
            }
          | {
              created_date?: string;
              text?: string;
            }
        )[];
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: (
          | {
              format_metadata?: {
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
          | {
              length?: number;
              start_index?: number;
              url_metadata?: {
                image_url?: string;
                snippet?: string;
                title?: string;
                url?: {
                  private_do_not_access_or_else_safe_url_wrapped_value?: string;
                };
              };
            }
        )[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          created_date?: string;
          quoted_message_metadata?: {
            annotations?: {
              length?: number;
              start_index?: number;
              youtube_metadata?: {
                id?: string;
                start_time?: number;
              };
            }[];
            creator?: {
              email?: string;
              name?: string;
              user_type?: string;
            };
            text?: string;
          };
          text?: string;
        }[];
        quoted_message_metadata?: {
          annotations?: {
            length?: number;
            start_index?: number;
            youtube_metadata?: {
              id?: string;
              start_time?: number;
            };
          }[];
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          url_metadata?: {
            image_url?: string;
            snippet?: string;
            title?: string;
            url?: {
              private_do_not_access_or_else_safe_url_wrapped_value?: string;
            };
          };
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          url_metadata?: {
            image_url?: string;
            snippet?: string;
            title?: string;
            url?: {
              private_do_not_access_or_else_safe_url_wrapped_value?: string;
            };
          };
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        topic_id?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          annotations?: {
            format_metadata?: {
              format_type?: string;
            };
            length?: number;
            start_index?: number;
          }[];
          attached_files?: {
            export_name?: string;
            original_name?: string;
          }[];
          created_date?: string;
          quoted_message_metadata?: {
            creator?: {
              email?: string;
              name?: string;
              user_type?: string;
            };
            text?: string;
          };
          text?: string;
        }[];
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          youtube_metadata?: {
            id?: string;
            start_time?: number;
          };
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: (
          | {
              annotations?: {
                length?: number;
                start_index?: number;
                youtube_metadata?: {
                  id?: string;
                  start_time?: number;
                };
              }[];
              text?: string;
              updated_date?: string;
            }
          | {
              annotations?: {
                length?: number;
                start_index?: number;
                youtube_metadata?: {
                  id?: string;
                  start_time?: number;
                };
              }[];
              created_date?: string;
              text?: string;
            }
        )[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          url_metadata?: {
            image_url?: string;
            snippet?: string;
            title?: string;
            url?: {
              private_do_not_access_or_else_safe_url_wrapped_value?: string;
            };
          };
        }[];
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          annotations?: (
            | {
                format_metadata?: {
                  format_type?: string;
                };
                length?: number;
                start_index?: number;
              }
            | {
                drive_metadata?: {
                  id?: string;
                  thumbnail_url?: string;
                  title?: string;
                };
                length?: number;
                start_index?: number;
              }
          )[];
          created_date?: string;
          text?: string;
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          annotations?: {
            format_metadata?: {
              format_type?: string;
            };
            length?: number;
            start_index?: number;
          }[];
          attached_files?: {
            export_name?: string;
            original_name?: string;
          }[];
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        topic_id?: string;
      }
    | {
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          annotations?: (
            | {
                format_metadata?: {
                  format_type?: string;
                };
                length?: number;
                start_index?: number;
              }
            | {
                length?: number;
                start_index?: number;
                url_metadata?: {
                  image_url?: string;
                  snippet?: string;
                  title?: string;
                  url?: {
                    private_do_not_access_or_else_safe_url_wrapped_value?: string;
                  };
                };
              }
          )[];
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: (
          | {
              length?: number;
              start_index?: number;
              url_metadata?: {
                image_url?: string;
                snippet?: string;
                title?: string;
                url?: {
                  private_do_not_access_or_else_safe_url_wrapped_value?: string;
                };
              };
            }
          | {
              length?: number;
              start_index?: number;
              youtube_metadata?: {
                id?: string;
                start_time?: number;
              };
            }
        )[];
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: (
          | {
              annotations?: {
                format_metadata?: {
                  format_type?: string;
                };
                length?: number;
                start_index?: number;
              }[];
              text?: string;
              updated_date?: string;
            }
          | {
              created_date?: string;
              text?: string;
            }
        )[];
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            font_color?: number;
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: (
          | {
              annotations?: {
                format_metadata?: {
                  format_type?: string;
                };
                length?: number;
                start_index?: number;
              }[];
              text?: string;
              updated_date?: string;
            }
          | {
              text?: string;
              updated_date?: string;
            }
          | {
              created_date?: string;
              text?: string;
            }
        )[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          annotations?: {
            length?: number;
            start_index?: number;
            youtube_metadata?: {
              id?: string;
              start_time?: number;
            };
          }[];
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          url_metadata?: {
            image_url?: string;
            snippet?: string;
            title?: string;
            url?: {
              private_do_not_access_or_else_safe_url_wrapped_value?: string;
            };
          };
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          annotations?: {
            length?: number;
            start_index?: number;
            url_metadata?: {
              image_url?: string;
              snippet?: string;
              title?: string;
              url?: {
                private_do_not_access_or_else_safe_url_wrapped_value?: string;
              };
            };
          }[];
          created_date?: string;
          text?: string;
        }[];
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: (
          | {
              format_metadata?: {
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
          | {
              length?: number;
              start_index?: number;
              url_metadata?: {
                image_url?: string;
                snippet?: string;
                title?: string;
                url?: {
                  private_do_not_access_or_else_safe_url_wrapped_value?: string;
                };
              };
            }
        )[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          annotations?: (
            | {
                format_metadata?: {
                  format_type?: string;
                };
                length?: number;
                start_index?: number;
              }
            | {
                length?: number;
                start_index?: number;
                url_metadata?: {
                  image_url?: string;
                  snippet?: string;
                  title?: string;
                  url?: {
                    private_do_not_access_or_else_safe_url_wrapped_value?: string;
                  };
                };
              }
          )[];
          created_date?: string;
          text?: string;
        }[];
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: (
          | {
              length?: number;
              start_index?: number;
              url_metadata?: {
                image_url?: string;
                snippet?: string;
                title?: string;
                url?: {
                  private_do_not_access_or_else_safe_url_wrapped_value?: string;
                };
              };
            }
          | {
              format_metadata?: {
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
        )[];
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: (
          | {
              format_metadata?: {
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
          | {
              format_metadata?: {
                font_color?: number;
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
        )[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          annotations?: (
            | {
                format_metadata?: {
                  format_type?: string;
                };
                length?: number;
                start_index?: number;
              }
            | {
                format_metadata?: {
                  font_color?: number;
                  format_type?: string;
                };
                length?: number;
                start_index?: number;
              }
          )[];
          created_date?: string;
          text?: string;
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          attached_files?: {
            export_name?: string;
            original_name?: string;
          }[];
          created_date?: string;
          quoted_message_metadata?: {
            creator?: {
              email?: string;
              name?: string;
              user_type?: string;
            };
            text?: string;
          };
          text?: string;
        }[];
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: (
          | {
              annotations?: {
                format_metadata?: {
                  format_type?: string;
                };
                length?: number;
                start_index?: number;
              }[];
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              text?: string;
              updated_date?: string;
            }
          | {
              annotations?: {
                format_metadata?: {
                  format_type?: string;
                };
                length?: number;
                start_index?: number;
              }[];
              attached_files?: {
                export_name?: string;
                original_name?: string;
              }[];
              created_date?: string;
              text?: string;
            }
        )[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          youtube_metadata?: {
            id?: string;
            start_time?: number;
          };
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          url_metadata?: {
            image_url?: string;
            snippet?: string;
            title?: string;
            url?: {
              private_do_not_access_or_else_safe_url_wrapped_value?: string;
            };
          };
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        attached_files?: {
          export_name?: string;
          original_name?: string;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          attached_files?: {
            export_name?: string;
            original_name?: string;
          }[];
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          url_metadata?: {
            image_url?: string;
            snippet?: string;
            title?: string;
            url?: {
              private_do_not_access_or_else_safe_url_wrapped_value?: string;
            };
          };
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          annotations?: {
            length?: number;
            start_index?: number;
            url_metadata?: {
              image_url?: string;
              snippet?: string;
              title?: string;
              url?: {
                private_do_not_access_or_else_safe_url_wrapped_value?: string;
              };
            };
          }[];
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
      }
    | {
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          annotations?: {
            format_metadata?: {
              format_type?: string;
            };
            length?: number;
            start_index?: number;
          }[];
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
      }
    | {
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          annotations?: {
            format_metadata?: {
              format_type?: string;
            };
            length?: number;
            start_index?: number;
          }[];
          attached_files?: {
            export_name?: string;
            original_name?: string;
          }[];
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: (
          | {
              format_metadata?: {
                font_color?: number;
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
          | {
              format_metadata?: {
                format_type?: string;
              };
              length?: number;
              start_index?: number;
            }
        )[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          created_date?: string;
          quoted_message_metadata?: {
            creator?: {
              email?: string;
              name?: string;
              user_type?: string;
            };
            text?: string;
          };
          text?: string;
        }[];
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            font_color?: number;
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        quoted_message_metadata?: {
          creator?: {
            email?: string;
            name?: string;
            user_type?: string;
          };
          text?: string;
        };
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          length?: number;
          start_index?: number;
          youtube_metadata?: {
            id?: string;
            start_time?: number;
          };
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: {
          annotations?: {
            length?: number;
            start_index?: number;
            youtube_metadata?: {
              id?: string;
              start_time?: number;
            };
          }[];
          created_date?: string;
          text?: string;
        }[];
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        annotations?: (
          | {
              length?: number;
              start_index?: number;
              url_metadata?: {
                image_url?: string;
                snippet?: string;
                title?: string;
                url?: {
                  private_do_not_access_or_else_safe_url_wrapped_value?: string;
                };
              };
            }
          | {
              length?: number;
              start_index?: number;
              youtube_metadata?: {
                id?: string;
                start_time?: number;
              };
            }
        )[];
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        text?: string;
        topic_id?: string;
      }
    | {
        annotations?: {
          format_metadata?: {
            format_type?: string;
          };
          length?: number;
          start_index?: number;
        }[];
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        message_id?: string;
        previous_message_versions?: (
          | {
              annotations?: {
                format_metadata?: {
                  format_type?: string;
                };
                length?: number;
                start_index?: number;
              }[];
              text?: string;
              updated_date?: string;
            }
          | {
              annotations?: {
                format_metadata?: {
                  format_type?: string;
                };
                length?: number;
                start_index?: number;
              }[];
              created_date?: string;
              text?: string;
            }
        )[];
        reactions?: {
          emoji?: {
            unicode?: string;
          };
          reactor_emails?: string[];
        }[];
        text?: string;
        topic_id?: string;
        updated_date?: string;
      }
    | {
        created_date?: string;
        creator?: {
          email?: string;
          name?: string;
          user_type?: string;
        };
        deleted_date?: string;
        deletion_metadata?: {
          deletion_type?: string;
        };
        message_id?: string;
        message_state?: string;
        text?: string;
        topic_id?: string;
      }
  )[];
}
